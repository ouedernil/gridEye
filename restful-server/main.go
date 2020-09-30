package main

import (
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	"log"
	"net/http"
	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")
	router := NewRouter()
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
