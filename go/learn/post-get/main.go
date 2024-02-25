package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("POST"))
		case "GET":
			w.Write([]byte("GET"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("alr started dumbass")
	http.ListenAndServe(":9000", nil)
}
