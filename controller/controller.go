package controller

import (
	"fmt"
	"gophercise-03-cyoa/model"
	"html/template"
	"log"
	"net/http"
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
	hdc := h.Dc
	if story, found := hdc.GetByPath(r.URL.Path); found {
		tmpl := template.Must(h.Tmpl, nil)
		err := tmpl.Execute(w, story)
		if (err != nil) {
			log.Fatalln(err)
		}
		return
	}
	if tryRedirect(w, r) {
		return
	}
	fmt.Fprint(w, "Error: Counld not find the entry point of the story.")
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
