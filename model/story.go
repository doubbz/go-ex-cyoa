package model

type Chapter struct {
	Name    string    `json:"name"`
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []options `json:"options"`
}
type options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Story map[string]Chapter
