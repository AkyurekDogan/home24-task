package htmlanalyzer

import (
	"context"

	"github.com/AkyurekDogan/home24-task/internal/app/model"

	"github.com/PuerkitoBio/goquery"
)

// Plugin is the interface that wraps the Do method for HTML analysis plugins.
type Plugin interface {
	Do(ctx context.Context, htmlDoc *goquery.Document, ar *model.AnalysisResult)
}
