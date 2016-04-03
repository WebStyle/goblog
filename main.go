package main

import (
	"net/http"
	"goblog/handlers"
)

func main() {

	// Routing
	http.HandleFunc("/", handlers.IndexHandler)
  	http.HandleFunc("/post/", handlers.PostCreateHandler)
	http.HandleFunc("/post/save", handlers.PostSaveHandler)
	http.HandleFunc("/about/", handlers.AboutHandler)

	// Static files
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Run a server
	http.ListenAndServe(":12345", nil)

}
