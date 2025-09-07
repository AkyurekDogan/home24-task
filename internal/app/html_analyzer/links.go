package htmlanalyzer

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type linksPlugin struct {
}

func NewLinksPlugin() Plugin {
	return &linksPlugin{}
}

func (vp *linksPlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult) {
	links := getLinks(htmlDoc)
	ar.TotalLinks = len(links)
	internal, external := classifyLinks(ar.URL, ar.FinalURL, links)
	ar.InternalLinks = internal
	ar.ExternalLinks = external
	url, _ := url.Parse(ar.URL)
	inaccessible := checkLinksAccessibility(ctx, url, links)
	ar.Inaccessible = inaccessible
}

func getLinks(doc *goquery.Document) []string {
	links := []string{}
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		h, _ := s.Attr("href")
		if h != "" {
			links = append(links, strings.TrimSpace(h))
		}
	})
	return links
}

func classifyLinks(baseHost, responseUrl string, links []string) (int, int) {
	// classify internal vs external
	var internal, external int
	for _, l := range links {
		// ignore anchors and mailto/tel/javascript
		if strings.HasPrefix(l, "#") || strings.HasPrefix(l, "mailto:") || strings.HasPrefix(l, "tel:") || strings.HasPrefix(l, "javascript:") {
			continue
		}
		abs, err := url.Parse(responseUrl)
		if err != nil {
			continue
		}
		if abs.Hostname() == "" || sameHostname(baseHost, abs.Hostname()) {
			internal++
		} else {
			external++
		}
	}
	return internal, external
}

func sameHostname(a, b string) bool {
	// simple compare without subdomain normalization
	return strings.EqualFold(a, b)
}

// checkLinksAccessibility checks links concurrently and returns count of inaccessible links.
func checkLinksAccessibility(ctx context.Context, base *url.URL, links []string) int {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10) // max 10 concurrent checks
	var mu sync.Mutex
	count := 0

	client := &http.Client{
		Timeout: 8 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // tolerate some cert issues for accessibility check
		},
	}

	for _, l := range links {
		ln := l
		// skip anchors and non-http(s)
		if strings.HasPrefix(ln, "#") || strings.HasPrefix(ln, "mailto:") || strings.HasPrefix(ln, "tel:") || strings.HasPrefix(ln, "javascript:") {
			continue
		}
		abs, err := base.Parse(ln)
		if err != nil {
			continue
		}
		if abs.Scheme != "http" && abs.Scheme != "https" {
			continue
		}
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			// try HEAD first
			req, _ := http.NewRequestWithContext(ctx, "HEAD", u, nil)
			req.Header.Set("User-Agent", "url-analyzer/1.0")
			resp, err := client.Do(req)
			if err != nil {
				// try GET as fallback
				req2, _ := http.NewRequestWithContext(ctx, "GET", u, nil)
				req2.Header.Set("User-Agent", "url-analyzer/1.0")
				resp2, err2 := client.Do(req2)
				if err2 != nil || resp2.StatusCode >= 400 {
					mu.Lock()
					count++
					mu.Unlock()
					return
				}
				if resp2 != nil {
					resp2.Body.Close()
				}
				return
			}
			if resp != nil {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 {
					mu.Lock()
					count++
					mu.Unlock()
					return
				}
			}
		}(abs.String())
	}

	wg.Wait()
	return count
}
