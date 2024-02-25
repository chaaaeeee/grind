package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("view.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))
		var err = r.ParseForm() // parsing data yang dikirim dari view / aksi dari / ke /process 

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// input data from submitted form
		var name = r.FormValue("name")
		var message = r.Form.Get("message") // akses method form dulu, setelah itu di GET

		var data = map[string]string{"name": name, "message": message} // tampung data
		
		// sisip data atau applies template to data
		err = tmpl.Execute(w, data) 
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}	

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/process", routeSubmitPost)

	fmt.Println("alr started dumbass")
	http.ListenAndServe(":9000", nil)
}
