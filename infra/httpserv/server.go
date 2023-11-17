package httpserv

import (
	"fmt"
	"gophercise-03-cyoa/model"
	"log"
	"net/http"
)

type Server struct {
	s StoryHandler
}

type StoryHandler struct {
	dc *dataCient
}

type dataCient interface {
	GetByPath(path string) (model.Chapter, bool)
}

func NewServer(dc dataCient) Server {
	return Server{StoryHandler{&dc}}
}

func (s Server) Serve() {
	mux := http.NewServeMux()
	mux.Handle("/", s.s)
	fmt.Println("listening")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
