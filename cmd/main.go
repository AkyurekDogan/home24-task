package main

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"

	"github.com/AkyurekDogan/home24-task/internal/app/handler"
	htmlanalyzer "github.com/AkyurekDogan/home24-task/internal/app/html_analyzer"
	"github.com/AkyurekDogan/home24-task/internal/app/infrastructure/config"
	"github.com/AkyurekDogan/home24-task/internal/app/infrastructure/logger"
	"github.com/AkyurekDogan/home24-task/internal/app/requester"
	"github.com/AkyurekDogan/home24-task/internal/app/service"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	// ENV environment file path
	ENV = ".env"
	//ENV_CNF_PATH config path
	ENV_CNF_PATH = "CONFIG_PATH"
)

func main() {
	// Initialize structured logger
	log, err := logger.NewLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer log.Sync()      // Flush any buffered log entries
	logger := log.Sugar() // for easy to use for simple projects.

	// load environment variables
	err = godotenv.Load(ENV)
	if err != nil {
		logger.Fatalf("Error loading .env file: " + err.Error())
	}
	// use environment variable to get the config path
	appEnvConfigPath := os.Getenv(ENV_CNF_PATH)
	if appEnvConfigPath == "" {
		logger.Fatalf("environment variable must be set: " + ENV_CNF_PATH)
	}
	// unmarshall the config file and get all settings in the configuration file.
	yamlFile, err := os.ReadFile(appEnvConfigPath)
	if err != nil {
		logger.Fatalf("error reading configuration YAML file: " + err.Error())
	}
	var config config.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		logger.Fatalf("error unmarshalling YAML file: " + err.Error())
	}
	// Render the templete
	tmpl := template.Must(template.ParseFiles(config.App.Template.Path))
	// create the http client with timeout and tls config
	// to avoid certificate issues breaking many requests when testing; still allow insecure if needed
	client := &http.Client{
		Timeout: config.App.Requester.Http.Timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
			DialContext: (&net.Dialer{
				Timeout:   config.App.Requester.Http.DialContext.Timeout,
				KeepAlive: config.App.Requester.Http.DialContext.KeepAlive,
			}).DialContext,
		},
	}
	requester := requester.New(client)
	htmlAnalyzer := htmlanalyzer.NewHTMLAnalyzer(
		htmlanalyzer.NewVersionPlugin(),
		htmlanalyzer.NewTitlePlugin(),
		htmlanalyzer.NewHeaderPlugin(),
		htmlanalyzer.NewLinksPlugin(),
		htmlanalyzer.NewLoginFormCheckerPlugin(),
	)

	analyzerService := service.NewAnalyzer(requester, htmlAnalyzer)
	handler := handler.NewAnalyzer(logger, tmpl, analyzerService)

	r := chi.NewRouter()

	// GET handler – render empty form
	r.Get("/", handler.Get)
	r.Post("/", handler.Analyze)
	// Reject any other HTTP methods on "/" with 405 Method Not Allowed
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed", r.Method)
	})
	addr := fmt.Sprintf("%s:%d",
		config.App.Service.Host,
		config.App.Service.Port,
	)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("Listening on %s", addr)
}
