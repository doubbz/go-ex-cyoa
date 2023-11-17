package main

import (
	"fmt"
	"gophercise-03-cyoa/cyoatypes"
	"html/template"
	"log"
	"net/http"
	"os"
)

type StoryHandler struct {
	s cyoatypes.Story
}

func (h StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if story, found := h.s[r.URL.Path]; found {
		tmpl := template.Must(template.ParseFiles("templates/layout.html"))
		tmpl.Execute(w, story)
		return
	}
	if _, found := h.s["/intro"]; found {
		http.Redirect(w, r, "/intro", http.StatusFound)
		return
	}
	fmt.Fprint(w, "Error: Counld not find the entry point of the story.")
}

func main() {
	file, err := os.Open("assets/story.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var h StoryHandler
	h.s = cyoatypes.JSONtoMap(file)

	mux := http.NewServeMux()
	mux.Handle("/", h)
	fmt.Println("listening")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
