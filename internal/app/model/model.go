package model

import "net/url"

const (
	UrlScopeInternal = "internal"
	UrlScopeExternal = "external"
)

type AnalysisResult struct {
	Query        Query
	Version      string
	Title        string
	Headers      map[string]int
	HasLoginForm bool
	Links        Links
}

type Query struct {
	BaseUrl url.URL
}

type Links []LinkAnalysis

func (l Links) Count() int {
	return len(l)
}
func (l Links) GetCounts(scope string, isAccessible bool) int {
	count := 0
	for _, link := range l {
		if link.Scope == scope {
			if link.IsAccessible == isAccessible {
				count++
			}
		}
	}
	return count
}

type LinkAnalysis struct {
	Url          url.URL
	Scope        string
	IsAccessible bool
}
