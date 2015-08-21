package handlers

import (
  "net/http"
  "html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  t := template.New("index")
  t, err := t.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		panic(err)
	}
  t.Execute(w, "index")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("post")
  t, err := t.ParseFiles("templates/post.html", "templates/header.html","templates/footer.html")
	if err != nil {
	  panic(err)
	}
	t.Execute(w, "post")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("about")
  t, err := t.ParseFiles("templates/about.html", "templates/header.html","templates/footer.html")
	if err != nil {
	  panic(err)
	}
	t.Execute(w, "about")
}
