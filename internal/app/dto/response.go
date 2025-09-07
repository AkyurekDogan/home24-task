package dto

type Analyze struct {
	URL           string
	Version       string
	Title         string
	HeaderCounts  map[string]int
	TotalLinks    int
	InternalLinks Link
	ExternalLinks Link
	HasLoginForm  string
	Error         string
}

type Link struct {
	Accessible   int
	Inaccessible int
}

type Error struct {
	Code    int
	Message string
}

type Response struct {
	Result *Analyze
	Error  *Error
}
