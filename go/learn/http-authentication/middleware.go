package main 

import "net/http"

const USERNAME = "batman"
const PASSWORD = "wayne"

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth() // ok ngecek valid engganya basic auth request
	if !ok {
		w.Write([]byte(`something went wrong`)) // error message kalo ok == false
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD) 
	if !isValid {
		w.Write([]byte(`wrong username or password`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("get only lol"))
	}
	
	return true
}
