package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The GodFactor", Director: "Francis Ford"},
				{Title: "Blade runner", Director: "Ridely scott"},
				{Title: "The thing", Director: "John Carpenter Ford"},
			},
		}
		templ.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 4)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})

	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film", h2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
