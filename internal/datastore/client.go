package datastore

import (
	"encoding/json"
	"gophercise-03-cyoa/internal/model"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

type Client struct {
	data model.Story
}

func NewClient() Client {
	dataPath := os.Getenv("APP_DATA_PATH")
	if (dataPath == "") {
		dataPath = "assets/story.json"
	}
	file, err := os.Open(dataPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error during data read.")
	}
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()
	
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
	err := jsonData.Decode(&c)
	if (err != nil) {
		log.Error().Err(err).Stack().Msg("Something went wrong while decoding the JSON.")
	}

	s := model.Story{}
	for i, v := range c {
		s["/"+c[i].Name] = v
	}

	return s
}
