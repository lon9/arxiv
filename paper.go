package arxiv

// Paper is object for paper
type Paper struct {
	ID          uint     `json:"id"`
	ArxivKey    string   `json:"arxivKey"`
	Title       string   `json:"title"`
	Authors     []Author `json:"authors"`
	Subject     string   `json:"subject"`
	Description string   `json:"description"`
}

// Author is object for author.
type Author struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	PaperID uint   `json:"paperId"`
}
