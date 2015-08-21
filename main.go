package main

import (
	"fmt"
	"net/http"
	"bot/models"
)

func main() {

	http.HandleFunc("/", handlers.IndexHandler)
  http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/about/", handlers.AboutHandler)
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.ListenAndServe(":12345", nil)
	fmt.Println("Server is run!")

}
