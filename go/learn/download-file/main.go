package main

import (
	"fmt"
	"net/http"
	"html/template"
	"path/filepath"
	"io"
	"os"
	"encoding/json"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)	
	}
}

func handleFilesList(w http.ResponseWriter, r *http.Request) {
	files := []M{}
	basePath, _ := os.Getwd()
	filesLocation := filepath.Join(basePath, "files")
	
	err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error) {
		if err != nil {
			return err
		}

		if info.IsDir() {

		}
	}
}

func main() {
	http.HandleFunc("/", handleIndex) 
	http.HandleFunc("/list-files", handleFilesList)
	http.HandleFunc("/download", handleDownload)

	fmt.Println("alr started dumbass")
	http.ListenAndServe(":9000", nil)
}
