package cyoatypes

import (
	"encoding/json"
	"io"
)

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

func JSONtoMap(r io.Reader) Story {
	jsonData := json.NewDecoder(r)
	var c []Chapter
	jsonData.Decode(&c)

	s := Story{}
	for i, v := range c {
		s["/"+c[i].Name] = v
	}

	return s
}
