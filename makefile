# Makefile

APP_NAME := url-analyzer
MAIN_FILE := ./cmd/main.go
ENV_FILE := .env
ENV_DIST := .env.dist
CONFIG_FILE := config.yml
CONFIG_DIST := config.yaml.dist
PORT := 8080

.PHONY: run

run:
	@echo "Setting up environment..."
	@cp -n $(ENV_DIST) $(ENV_FILE) 2>/dev/null || true
	@cp -n $(CONFIG_DIST) $(CONFIG_FILE) 2>/dev/null || true
	@echo "Running the application..."
	@go run $(MAIN_FILE) & \
	sleep 2; \
	if which open > /dev/null; then \
		open http://localhost:$(PORT); \
	elif which xdg-open > /dev/null; then \
		xdg-open http://localhost:$(PORT); \
	else \
		echo "Open your browser at http://localhost:$(PORT)"; \
	fi
