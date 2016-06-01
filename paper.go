package arxiv

// Paper is object for paper
type Paper struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Authors     []Author `json:"authors"`
	Subject     string   `json:"subject"`
	Description string   `json:"description"`
}

// Author is object for author.
type Author struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
