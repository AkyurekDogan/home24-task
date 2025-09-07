package htmlanalyzer

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	httpclient "github.com/AkyurekDogan/home24-task/internal/app/infrastructure/http_client"
	"github.com/AkyurekDogan/home24-task/internal/app/model"
	"github.com/PuerkitoBio/goquery"
)

type linksPlugin struct {
	httpClient httpclient.Doer
}

// NewLinksPlugin creates a new instance of linksPlugin.
func NewLinksPlugin(httpClient httpclient.Doer) Plugin {
	return &linksPlugin{
		httpClient: httpClient,
	}
}

// Do analyzes the HTML document for links and updates the AnalysisResult.
func (lp *linksPlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *model.AnalysisResult) {
	// get all links
	urls := lp.getLinks(ar.Query, htmlDoc) // only valid urls.
	for _, u := range urls {
		linkAnalysis := model.LinkAnalysis{Url: u}
		linkAnalysis.Scope = lp.classifyLink(ar.Query, u)
		linkAnalysis.IsAccessible = lp.isLinkAccessible(ctx, u)
		ar.Links = append(ar.Links, linkAnalysis)
	}
}

func (lp *linksPlugin) getLinks(query model.Query, doc *goquery.Document) []url.URL {
	links := []url.URL{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		h, _ := s.Attr("href")
		if h != "" {
			normalizedUrl, err := lp.NormalizeLink(&query.BaseUrl, h)
			if err == nil || normalizedUrl != nil { // only valid links and urls.
				links = append(links, *normalizedUrl)
			}
		}
	})
	return links
}

// NormalizeLink resolves a raw href (relative or absolute) against the base URL.
// Returns the absolute URL and its hostname.
func (lp *linksPlugin) NormalizeLink(base *url.URL, rawHref string) (*url.URL, error) {
	rawHref = strings.TrimSpace(rawHref)
	if rawHref == "" {
		return nil, nil // ignore empty links
	}
	// Parse the raw link
	parsed, err := url.Parse(rawHref)
	if err != nil {
		return nil, err
	}
	// Resolve relative URLs against the base
	normalized := base.ResolveReference(parsed)
	return normalized, nil
}

func (lp *linksPlugin) classifyLink(query model.Query, link url.URL) string {
	if lp.sameHostname(query.BaseUrl.Host, link.Hostname()) {
		return model.UrlScopeInternal
	}
	return model.UrlScopeExternal
}

func (lp *linksPlugin) sameHostname(a, b string) bool {
	// simple compare without subdomain normalization
	return strings.EqualFold(a, b)
}

// isLinkAccessible checks a single URL and returns true if it is accessible, false otherwise.
func (lp *linksPlugin) isLinkAccessible(ctx context.Context, link url.URL) bool {
	// Only HTTP/HTTPS links are relevant
	if link.Scheme != "http" && link.Scheme != "https" {
		return false
	}
	// Try HEAD first since it is cheaper.
	req, _ := http.NewRequestWithContext(ctx, "HEAD", link.String(), nil)
	req.Header.Set("User-Agent", "url-analyzer/1.0")
	resp, err := lp.httpClient.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		// fallback to GET if HEAD fails
		req2, _ := http.NewRequestWithContext(ctx, "GET", link.String(), nil)
		req2.Header.Set("User-Agent", "url-analyzer/1.0")
		resp2, err2 := lp.httpClient.Do(req2)
		if err2 != nil || resp2.StatusCode >= 400 {
			return false
		}
		if resp2 != nil {
			resp2.Body.Close()
		}
		return true
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	return true
}
