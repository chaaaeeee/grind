package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	// ngebuat data dalam bentuk array of struct
	// struct isi var nya string dan int dijadikan array jadi bisa lebih dari 1 data
	data := [] struct {
		Name string
		Age int
	} {
		{"Richard Grayson", 24 },
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne ", 21},
	}
	
	// mereturn json encoding dari data dalam bentuk slices of bytes
	// merubah data structure dari n ke json string, pada kasus ini data bertipe slices of structs
	
	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	// ngasih content type, hapus aja, ntar keliatan bedanya
	w.Header().Set("Content-Type", "application/json")
	
	w.Write(jsonInBytes)

	// ngemarshal sekalian ngewrite deh kayanya
	/*
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	*/	
}

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("alr started dumbass")
	http.ListenAndServe(":9000", nil)
}
