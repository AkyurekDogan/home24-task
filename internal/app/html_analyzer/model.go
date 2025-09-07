package htmlanalyzer

type AnalysisResult struct {
	URL            string
	FinalURL       string // after redirects
	HTMLVersion    string
	Title          string
	HeadingCounts  map[string]int
	TotalLinks     int
	InternalLinks  int
	ExternalLinks  int
	Inaccessible   int
	LoginFormFound bool
}
