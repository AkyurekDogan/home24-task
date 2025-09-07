package htmlanalyzer

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type versionPlugin struct {
}

func NewVersionPlugin() Plugin {
	return &versionPlugin{}
}

func (vp *versionPlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult) {
	ar.HTMLVersion = getHTMLVersion(htmlDoc)
}

func getHTMLVersion(doc *goquery.Document) string {
	for _, n := range doc.Nodes {
		if n.Type == html.DocumentNode {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.DoctypeNode {
					doctype := strings.ToLower(c.Data)
					if strings.Contains(doctype, "html") && !strings.Contains(doctype, "xhtml") {
						return "HTML5 (<!DOCTYPE html>)"
					}
					if strings.Contains(doctype, "xhtml") {
						return "XHTML"
					}
					if strings.Contains(doctype, "html public") || strings.Contains(doctype, "dtd html") {
						return "HTML 4.x (DTD)"
					}
					break
				}
			}
		}
	}
	return "Unknown/other"
}
