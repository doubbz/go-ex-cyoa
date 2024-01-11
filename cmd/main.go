package cmd

import (
	"fmt"
	"gophercise-03-cyoa/internal/controller"
	"gophercise-03-cyoa/internal/datastore"
	"html/template"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func init() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

func Execute() {

	tmplPath := os.Getenv("APP_TEMPLATE_PATH") 
	if tmplPath == "" {
		tmplPath = "./templates/layout.html"
	}
	
	client := datastore.NewClient()
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("Something went wrong while parsing the template.")
	}

	port := os.Getenv("APP_PORT") 
	if port == "" {
		port = "8080"
	}

	h := controller.NewHandler(client, tmpl)
	mux := http.NewServeMux()
	mux.Handle("/", h)
	
	logger.Debug().Msg("Application has booted successfully.")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	
	if err != nil {
		logger.Fatal().Err(err).Msg("Something went wrong with server.")
	}
}
