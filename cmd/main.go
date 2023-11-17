package cmd

import (
	"gophercise-03-cyoa/infra/datastore"
	"gophercise-03-cyoa/infra/httpserv"
)

func Execute() {
	client := datastore.NewClient()
	server := httpserv.NewServer(client)
	server.Serve()
}
