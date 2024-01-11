package controller

import (
	"gophercise-03-cyoa/model"
	"html/template"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

const storyIntro = "/intro"

type StoryHandler struct {
	Dc dataClient
	Tmpl *template.Template
}

type dataClient interface {
	GetByPath(path string) (model.Chapter, bool)
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mu := sync.Mutex{}
	defer mu.Unlock()
	
	resultChan := make(chan model.Chapter)
	go func (resultChan chan model.Chapter) {
		hdc := h.Dc
		if story, found := hdc.GetByPath(r.URL.Path); found {
			resultChan <- story
			return
		}
		resultChan <- model.Chapter{}
		}(resultChan)
	
	// For academic purposes, we will use the context provided by the net/http.Request
	select {
	case story := <- resultChan:
		mu.Lock()
		if story.Name != "" {
			tmpl := template.Must(h.Tmpl, nil)
			err := tmpl.Execute(w, story)
			if err != nil {
				log.Error().Err(err).Stack().Msg("Something went wrong while executing the template.")
				http.Error(w, "Oops! Something went wrong. Please try again or contact support.", http.StatusInternalServerError)
			}
			break
		}

		if tryRedirect(w, r) {
			break
		}
		
		log.Error().Stack().Msg("Could not find the entry point of the story.")
		http.Error(w, "Oops! Could not find the entry point of the story.", http.StatusNotFound)
	case <-ctx.Done():
		mu.Lock()
		err := ctx.Err()
		log.Error().Err(err).Stack().Msg("Context has been cancelled")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		break
	}
}

func tryRedirect(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != storyIntro {
		http.Redirect(w, r, storyIntro, http.StatusFound)
		return true
	}
	return false
}

func NewHandler(dc dataClient, tmpl *template.Template) StoryHandler {
	return StoryHandler{dc, tmpl}
}
