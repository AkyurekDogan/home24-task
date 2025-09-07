package logger

import (
	"go.uber.org/zap"
)

// NewLogger initializes and returns a new Zap logger instance.
func NewLogger() (*zap.Logger, error) {
	// Use production configuration for structured logging
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
