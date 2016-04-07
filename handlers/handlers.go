package handlers

import (
	"fmt"
  	"net/http"
  	"html/template"

	"goblog/models"
)

// Index Page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  	t := template.New("index")
  	t, err := t.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
  	t.Execute(w, "index")
}

// Post create page
func PostCreateHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("postCreate")
  	t, err := t.ParseFiles("templates/post.html", "templates/header.html","templates/footer.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "postCreate")
}

func PostSaveHandler(w http.ResponseWriter, r *http.Request) {

	title 	:= r.FormValue("title")
	content := r.FormValue("content")

	result := Post.NewPost(title, content)
	fmt.Println(result)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("about")
  	t, err := t.ParseFiles("templates/about.html", "templates/header.html","templates/footer.html")
	if err != nil {
	  panic(err)
	}
	t.Execute(w, "about")
}
