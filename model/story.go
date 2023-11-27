package model

type Chapter struct {
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}
type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Chapter
