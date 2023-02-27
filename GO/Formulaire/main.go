package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl1 := template.Must(template.ParseFiles("formulaire.html"))
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl1.Execute(w, nil)
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")
		fmt.Printf("Nom: %s\nEmail: %s\nMessage: %s\n", name, email, message)
	})
	http.ListenAndServe(":8080", nil)
}
