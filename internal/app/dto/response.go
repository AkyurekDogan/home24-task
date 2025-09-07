package dto

type Analyze struct {
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
	Error          string
}

type Error struct {
	Code    int
	Message string
}

type Response struct {
	Result *Analyze
	Error  *Error
}
