package service

import (
	"context"
	"fmt"

	htmlanalyzer "github.com/AkyurekDogan/home24-task/internal/app/html_analyzer"
	"github.com/AkyurekDogan/home24-task/internal/app/requester"
)

type Analyzer interface {
	Analyze(ctx context.Context, urlStr string) (*htmlanalyzer.AnalysisResult, error)
}

type analyzer struct {
	requesterService requester.Requester
	htmlAnalyzer     htmlanalyzer.HTMLAnalyzer
}

func NewAnalyzer(
	requesterService requester.Requester,
	htmlAnalyzer htmlanalyzer.HTMLAnalyzer,
) Analyzer {
	return &analyzer{
		requesterService: requesterService,
		htmlAnalyzer:     htmlAnalyzer,
	}
}

func (a *analyzer) Analyze(ctx context.Context, urlStr string) (*htmlanalyzer.AnalysisResult, error) {
	// Implement the analysis logic here.
	htmlDoc, err := a.requesterService.Do(ctx, urlStr)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML document: %w", err)
	}
	// Analyze the HTML document.
	result, err := a.htmlAnalyzer.Analyze(ctx, htmlDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze HTML document: %w", err)
	}
	return &result, nil
}
