package htmlanalyzer

import (
	"context"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type HTMLAnalyzer interface {
	Analyze(ctx context.Context, htmlDoc *goquery.Document) (AnalysisResult, error)
}
type htmlAnalyzer struct {
	plugins []Plugin
}

func NewHTMLAnalyzer(plugins ...Plugin) HTMLAnalyzer {
	return &htmlAnalyzer{
		plugins: plugins,
	}
}

func (ha *htmlAnalyzer) Analyze(ctx context.Context, htmlDoc *goquery.Document) (AnalysisResult, error) {
	var (
		result AnalysisResult
		mu     sync.Mutex // protects result (if multiple plugins write to it)
		wg     sync.WaitGroup
	)
	for _, p := range ha.plugins {
		wg.Add(1)
		go func(plugin Plugin) {
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
