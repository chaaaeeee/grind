package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type M map[string]interface{}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{
			"name": "Wayne",
			"known": "Batman",
			"wealth": "Unlimited",
			"status": "Unknown",
		}
		var tmpl = template.Must(template.ParseFiles (
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		
		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{
			"name": "Wayne",
			"known": "Batman",
			"wealth": "Unlimited",
			"status": "Unknown",
		}
		var tmpl = template.Must(template.ParseFiles (
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("alr started")
	http.ListenAndServe(":9000", nil)
}
