package main

import "net/http"
import "html/template"
import "fmt"
import "path/filepath"
import "io"
import "os"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "post only lol", http.StatusBadRequest)
		return
	}

	basePath, _ := os.Getwd()
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
	}

	w.Write([]byte(`all files uploaded`))

}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)

	fmt.Println("alr started dumbass")
	http.ListenAndServe(":9000", nil)
}
