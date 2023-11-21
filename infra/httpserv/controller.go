package httpserv

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const storyIntro = "/intro"

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hdc := h.Dc
	if story, found := hdc.GetByPath(r.URL.Path); found {
		tmpl := template.Must(template.ParseFiles("templates/layout.html"))
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
