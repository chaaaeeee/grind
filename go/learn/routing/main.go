package main

import "fmt"
import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("assets")))
	fmt.Println("alr started")
	http.ListenAndServe(":9000", nil)
}
