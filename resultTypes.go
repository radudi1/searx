package searx

// SearchResult is a single search result from the searx response
type SearchResult struct {
	Title     string   `json:"title"`
	Url       string   `json:"url"`
	Content   string   `json:"content"`
	Engine    string   `json:"engine"`
	Engines   []string `json:"engines"`
	Positions []int    `json:"positions"`
	Score     float64  `json:"score"`
	Category  string   `json:"category"`
}

// SearchResponse is the response returned by searx
type SearchResponse struct {
	Query           string         `json:"query"`
	Results         []SearchResult `json:"results"`
	NumberOfResults int            `json:"numberOfResults"`
}
