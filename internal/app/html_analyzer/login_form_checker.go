package htmlanalyzer

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type loginFormCheckerPlugin struct {
}

func NewLoginFormCheckerPlugin() Plugin {
	return &loginFormCheckerPlugin{}
}

func (vp *loginFormCheckerPlugin) Do(ctx context.Context, htmlDoc *goquery.Document, ar *AnalysisResult) {
	ar.LoginFormFound = detectLoginForm(htmlDoc)
}

func detectLoginForm(doc *goquery.Document) bool {
	found := false
	doc.Find("form").EachWithBreak(func(i int, s *goquery.Selection) bool {
		// look for password inputs
		if s.Find("input[type='password']").Length() > 0 {
			found = true
			return false
		}
		// look for password-like names or ids
		pwdInputs := s.Find("input")
		pwdInputs.EachWithBreak(func(i int, in *goquery.Selection) bool {
			name, _ := in.Attr("name")
			id, _ := in.Attr("id")
			typ, _ := in.Attr("type")
			lower := strings.ToLower(name + " " + id + " " + typ)
			if strings.Contains(lower, "password") || strings.Contains(lower, "pass") {
				found = true
				return false
			}
			return true
		})
		if found {
			return false
		}
		// detect submit with login text
		if s.Find("button").FilterFunction(func(i int, b *goquery.Selection) bool {
			text := strings.ToLower(strings.TrimSpace(b.Text()))
			return strings.Contains(text, "login") || strings.Contains(text, "sign in")
		}).Length() > 0 {
			found = true
			return false
		}
		return true
	})
	return found
}
