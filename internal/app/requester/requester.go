package requester

import (
	"context"
	"fmt"
	"net/http"

	httpclient "github.com/AkyurekDogan/home24-task/internal/app/infrastructure/http_client"
	"github.com/PuerkitoBio/goquery"
)

type Requester interface {
	Do(ctx context.Context, url string) (*goquery.Document, error)
}

type requester struct {
	httpClient httpclient.Doer
}

func New(httpClient httpclient.Doer) Requester {
	return &requester{
		httpClient: httpClient,
	}
}
func (r *requester) Do(ctx context.Context, url string) (*goquery.Document, error) {
	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "url-analyzer/1.0 (+https://example)")
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request is failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("remote returned status %d", resp.StatusCode)
	}
	// Use goquery to parse document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html document parse failed: %w", err)
	}
	return doc, nil
}
