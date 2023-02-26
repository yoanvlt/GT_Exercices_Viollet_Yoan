package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(13) + 3
	y := rand.Intn(13) + 3
	tableau := make([][]int, y)
	for i := range tableau {
		tableau[i] = make([]int, x)
		for j := range tableau[i] {
			tableau[i][j] = rand.Intn(100) + 1
		}
	}
	tmpl, err := template.ParseFiles("tableau.html")
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, tableau)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
