package main

import (
	"fmt"
	"net/http"
	"html/template"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "learning golang web",
			"name": "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	
	http.Handle("/static/",
			http.StripPrefix("/static/",
				http.FileServer(http.Dir("assets"))))

	fmt.Println("alr started")
	http.ListenAndServe(":9000", nil)
}
