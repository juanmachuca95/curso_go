package main

import (
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

	fs := http.FileServer(http.Dir("./views/assets/"))
	http.Handle("/views/assets/", http.StripPrefix("/views/assets/", fs))

	tmpl := template.Must(template.ParseFiles("views/layout.html"))

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
