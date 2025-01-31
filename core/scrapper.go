package core

type Job struct {
	Title       string `json:"title" csv:"title"`
	Company     string `json:"company" csv:"company"`
	Location    string `json:"location" csv:"location"`
	URL         string `json:"url" csv:"url"`
	Description string `json:"description" csv:"description"`
}

type Scraper interface {
	Name() string
	Scrape() ([]Job, error)
}
