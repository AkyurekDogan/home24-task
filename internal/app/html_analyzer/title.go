package htmlanalyzer

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type titlePlugin struct {
}

func NewTitlePlugin() Plugin {
	return &titlePlugin{}
}

func (vp *titlePlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult) {
	ar.Title = getTitle(htmlDoc)
}

func getTitle(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find("title").First().Text())
}
