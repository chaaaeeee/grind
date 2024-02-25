package main 

import "net/http"

const USERNAME = "batman"
const PASSWORD = "wayne"

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares{
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth() 
		if !ok {
			w.Write([]byte(`something went wrong`)) // error message kalo ok == false
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD) 
		if !isValid {
			w.Write([]byte(`wrong username or password`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGET(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("get only lol"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
