package htmlanalyzer

import (
	"context"
	"net/url"
	"sync"

	plugin "github.com/AkyurekDogan/home24-task/internal/app/html_analyzer/plugin"
	"github.com/AkyurekDogan/home24-task/internal/app/model"

	"github.com/PuerkitoBio/goquery"
)

type HTMLAnalyzer interface {
	Analyze(
		ctx context.Context,
		url url.URL,
		htmlDoc *goquery.Document,
	) (model.AnalysisResult, error)
}
type htmlAnalyzer struct {
	plugins []plugin.Plugin
}

func New(plugins ...plugin.Plugin) HTMLAnalyzer {
	return &htmlAnalyzer{
		plugins: plugins,
	}
}

func (ha *htmlAnalyzer) Analyze(
	ctx context.Context,
	url url.URL,
	htmlDoc *goquery.Document,
) (model.AnalysisResult, error) {
	var (
		result model.AnalysisResult
		mu     sync.Mutex // protects result (if multiple plugins write to it)
		wg     sync.WaitGroup
	)
	result.Query.BaseUrl = url
	for _, p := range ha.plugins {
		wg.Add(1)
		go func(plugin plugin.Plugin) {
			defer wg.Done()
			// Each plugin modifies result for lock to prevent data races
			mu.Lock()
			plugin.Do(ctx, htmlDoc, &result) // they are NOT using the same fields concurrently
			mu.Unlock()
		}(p)
	}
	wg.Wait() // wait for all plugins to finish
	return result, nil
}
