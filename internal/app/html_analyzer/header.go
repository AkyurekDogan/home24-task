package htmlanalyzer

import (
	"context"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type headerPlugin struct {
}

func NewHeaderPlugin() Plugin {
	return &headerPlugin{}
}

func (vp *headerPlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult) {
	ar.HeadingCounts = getHeaders(htmlDoc)
}

func getHeaders(doc *goquery.Document) map[string]int {
	result := make(map[string]int)
	for i := 1; i <= 6; i++ { // total we have h1 to h6
		tag := fmt.Sprintf("h%d", i)
		result[tag] = doc.Find(tag).Length()
	}
	return result
}
