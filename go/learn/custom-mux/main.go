package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SelectStudent(id))
		return
	}
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}	

func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)
	
	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGET)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux 

	fmt.Println("alr started dumbass")
	server.ListenAndServe()
}
