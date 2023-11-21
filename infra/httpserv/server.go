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
	Dc dataClient
}

type dataClient interface {
	GetByPath(path string) (model.Chapter, bool)
}

func NewServer(dc dataClient) Server {
	return Server{StoryHandler{dc}}
}

func (s Server) Serve() {
	mux := http.NewServeMux()
	mux.Handle("/", s.s)
	fmt.Println("listening")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
