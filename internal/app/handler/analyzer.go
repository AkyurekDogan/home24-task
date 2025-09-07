package handler

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/AkyurekDogan/home24-task/internal/app/dto"
	htmlanalyzer "github.com/AkyurekDogan/home24-task/internal/app/html_analyzer"
	"github.com/AkyurekDogan/home24-task/internal/app/service"

	"go.uber.org/zap"
)

type Analyzer interface {
	Get(w http.ResponseWriter, r *http.Request)
	Analyze(w http.ResponseWriter, r *http.Request)
}

type analyzer struct {
	logger          *zap.SugaredLogger
	template        *template.Template
	analyzerService service.Analyzer
}

func NewAnalyzer(
	logger *zap.SugaredLogger,
	template *template.Template,
	analyzerService service.Analyzer,
) Analyzer {
	return &analyzer{
		logger:          logger,
		template:        template,
		analyzerService: analyzerService,
	}
}

func (a *analyzer) Get(w http.ResponseWriter, r *http.Request) {
	err := a.template.Execute(w, nil)
	if err != nil {
		a.logger.Errorf("template execution failed for GET: %v", err)
		a.toResponseError(
			w,
			http.StatusBadRequest,
			"somethign went wrong please contact your support!",
		)
		return
	}
	a.logger.Info(":GET endpoint is called") // it should be debug in real cases
}

func (a *analyzer) Analyze(w http.ResponseWriter, r *http.Request) {
	// Get the form value.
	ctx := r.Context()
	urlStr := r.FormValue("url")
	urlSenitized, err := a.validateURL(urlStr)
	if err != nil {
		a.logger.Warnf("invalid url: %s, error: %+v", urlStr, err)
		a.toResponseError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf("invalid url: %s", err.Error()),
		)
		return
	}
	result, err := a.analyzerService.Analyze(ctx, urlSenitized)
	if err != nil {
		a.logger.Errorf("analyzer service is failed, error: %v", err)
		a.toResponseError(
			w,
			http.StatusInternalServerError,
			"somethign went wrong please contact your support!",
		)
		return
	}
	// Render the template with results
	a.toResponseSuccess(w, *result)
	a.logger.Infof(":POST endpoint is called with url: %s", urlSenitized) // it should be debug in real cases
}

func (a *analyzer) toResponseSuccess(
	w http.ResponseWriter,
	result htmlanalyzer.AnalysisResult,
) {
	response := dto.Response{
		Result: &dto.Analyze{
			URL:            result.URL,
			FinalURL:       result.FinalURL,
			HTMLVersion:    result.HTMLVersion,
			Title:          result.Title,
			HeadingCounts:  result.HeadingCounts,
			TotalLinks:     result.TotalLinks,
			InternalLinks:  result.InternalLinks,
			ExternalLinks:  result.ExternalLinks,
			Inaccessible:   result.Inaccessible,
			LoginFormFound: result.LoginFormFound,
		},
		Error: nil,
	}
	a.renderResponse(w, response)
}

func (a *analyzer) toResponseError(
	w http.ResponseWriter,
	code int,
	msg string,
) dto.Response {
	response := dto.Response{
		Result: nil,
		Error: &dto.Error{
			Code:    code,
			Message: msg,
		},
	}
	a.renderResponse(w, response)
	return response
}

func (a *analyzer) renderResponse(w http.ResponseWriter, response dto.Response) {
	if err := a.template.Execute(w, response); err != nil {
		a.logger.Errorf("template execution failed: %v", err)
		http.Error(
			w,
			"something went wrong, please contact your support!",
			http.StatusInternalServerError,
		)
	}
}

// ValidateURL ensures the URL is well-formed and uses http/https schema
func (a *analyzer) validateURL(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", errors.New("URL cannot be empty")
	}

	// Parse and check for errors
	parsed, err := url.ParseRequestURI(input)
	if err != nil {
		return "", errors.New("invalid URL format")
	}

	// Require http or https scheme
	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", errors.New("URL must start with http:// or https://")
	}

	// Require a host (domain or IP)
	if parsed.Host == "" {
		return "", errors.New("URL must have a valid host")
	}

	return parsed.String(), nil
}
