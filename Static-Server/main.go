package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name       string
	Occupation string
}

type ViewData struct {
	Title string
	Users []User
}

func main() {
	fs := http.FileServer(http.Dir("public"))

	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("templates/layout.html"))

		data := ViewData{
			Title: "User List",
			Users: []User{
				{Name: "John Doe", Occupation: "gardener"},
				{Name: "Roger Roe", Occupation: "driver"},
				{Name: "Thomas Green", Occupation: "teacher"},
			},
		}


		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
