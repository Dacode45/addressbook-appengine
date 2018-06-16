package main

import (
	"net/http"

	"google.golang.org/appengine"

	"github.com/Dacode45/addressbook/server"
)

func init() {
	server.SetupDB()
}

func main() {
	http.Handle("/", server.SetupRouter())
	appengine.Main()
}
