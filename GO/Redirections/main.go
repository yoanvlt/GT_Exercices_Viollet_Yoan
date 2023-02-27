package main

import (
	"net/http"
	"text/template"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl := template.Must(template.ParseFiles("404.html"))
		tmpl.Execute(w, nil)
	}
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))
	tmpl1 := template.Must(template.ParseFiles("page1.html"))
	tmpl2 := template.Must(template.ParseFiles("page2.html"))
	tmpl3 := template.Must(template.ParseFiles("page3.html"))
	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page1" {
			errorHandler(w, r, http.StatusNotFound)
			return
		} else {
			tmpl1.Execute(w, nil)
		}
	})
	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page2" {
			errorHandler(w, r, http.StatusNotFound)
			return
		} else {
			tmpl2.Execute(w, nil)
		}
	})
	http.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/page3" {
			errorHandler(w, r, http.StatusNotFound)
			return
		} else {
			tmpl3.Execute(w, nil)
		}
	})
	http.ListenAndServe(":8080", nil)
}
