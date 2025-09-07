package htmlanalyzer

import (
	"context"
	"fmt"

	"github.com/AkyurekDogan/home24-task/internal/app/model"
	"github.com/PuerkitoBio/goquery"
)

type headerPlugin struct {
}

// NewHeaderPlugin creates a new instance of headerPlugin.
func NewHeaderPlugin() Plugin {
	return &headerPlugin{}
}

// Do analyzes the HTML document and counts the occurrences of each header tag (h1 to h6).
func (hp *headerPlugin) Do(
	ctx context.Context,
	htmlDoc *goquery.Document,
	ar *model.AnalysisResult,
) {
	ar.Headers = getHeaders(htmlDoc)
}

func getHeaders(doc *goquery.Document) map[string]int {
	result := make(map[string]int)
	for i := 1; i <= 6; i++ { // total we have h1 to h6 limitedly
		tag := fmt.Sprintf("h%d", i)
		result[tag] = doc.Find(tag).Length()
	}
	return result
}
