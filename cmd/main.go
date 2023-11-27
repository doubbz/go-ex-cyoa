package cmd

import (
	"fmt"
	"gophercise-03-cyoa/controller"
	"gophercise-03-cyoa/datastore"
	"html/template"
	"log"
	"net/http"
)

func Execute() {
	client := datastore.NewClient()
	tmpl, err := template.ParseFiles("./templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}

	h := controller.NewHandler(client, tmpl)

	mux := http.NewServeMux()
	mux.Handle("/", h)
	fmt.Println("listening")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
