package htmlanalyzer

import (
	"context"
	"strings"

	"github.com/AkyurekDogan/home24-task/internal/app/model"
	"github.com/PuerkitoBio/goquery"
)

type titlePlugin struct {
}

// NewTitlePlugin creates a new instance of titlePlugin.
func NewTitlePlugin() Plugin {
	return &titlePlugin{}
}

// Do extracts the title from the HTML document and updates the AnalysisResult.
func (vp *titlePlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *model.AnalysisResult) {
	ar.Title = getTitle(htmlDoc)
}

func getTitle(doc *goquery.Document) string {
	return strings.TrimSpace(doc.Find("title").First().Text())
}
