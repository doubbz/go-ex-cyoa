package datastore

import (
	"encoding/json"
	"gophercise-03-cyoa/model"
	"io"
	"log"
	"os"
)

type Client struct {
	data model.Story
}

func NewClient() Client {
	file, err := os.Open("assets/story.json")
	
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	s := loadInMemory(file)

	return Client{s}
}

func (c Client) GetByPath(path string) (model.Chapter, bool) {
	v, found := c.data[path]
	
	return v, found
}

func loadInMemory(r io.Reader) model.Story {
	jsonData := json.NewDecoder(r)
	var c []model.Chapter
	jsonData.Decode(&c)

	s := model.Story{}
	for i, v := range c {
		s["/"+c[i].Name] = v
	}

	return s
}
