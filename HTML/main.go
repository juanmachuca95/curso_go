package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
	Link  string
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {

	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	tmpl, err := template.ParseFiles(
		"templates/home/index.html",
		"templates/globals/footer.html",
		"templates/globals/header.html",
	)

	if err != nil {
		log.Fatalf("Error al complicar home : %s", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "By Juan Machuca",
			Todos: []Todo{
				{Title: "User", Done: false, Link: "/user"},
				{Title: "Hello World", Done: true, Link: "hello"},
				{Title: "About Us", Done: true, Link: "about"},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
