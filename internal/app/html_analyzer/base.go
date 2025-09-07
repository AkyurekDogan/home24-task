package htmlanalyzer

import (
	"context"

	"github.com/PuerkitoBio/goquery"
)

type Plugin interface {
	Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult)
}
