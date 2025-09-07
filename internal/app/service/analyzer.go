package service

import (
	"context"
	"fmt"
	"net/url"

	htmlanalyzer "github.com/AkyurekDogan/home24-task/internal/app/html_analyzer"
	"github.com/AkyurekDogan/home24-task/internal/app/model"
	"github.com/AkyurekDogan/home24-task/internal/app/requester"
)

type Analyzer interface {
	Analyze(
		ctx context.Context,
		url url.URL,
	) (*model.AnalysisResult, error)
}

type analyzer struct {
	requesterService    requester.Requester
	htmlAnalyzerService htmlanalyzer.HTMLAnalyzer
}

func NewAnalyzer(
	requesterService requester.Requester,
	htmlAnalyzerService htmlanalyzer.HTMLAnalyzer,
) Analyzer {
	return &analyzer{
		requesterService:    requesterService,
		htmlAnalyzerService: htmlAnalyzerService,
	}
}

func (a *analyzer) Analyze(
	ctx context.Context,
	url url.URL,
) (*model.AnalysisResult, error) {
	// Implement the analysis logic here.
	htmlDoc, err := a.requesterService.Do(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML document: %w", err)
	}
	// Analyze the HTML document.
	result, err := a.htmlAnalyzerService.Analyze(ctx, url, htmlDoc)
	if err != nil {
		return nil, fmt.Errorf("failed to analyze HTML document: %w", err)
	}
	return &result, nil
}
