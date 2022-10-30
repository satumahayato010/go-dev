package main

import (
	"html/template"
	"log"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	file, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("no such HTMLFile")
	}
	file.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", HelloWorldHandler)

	http.ListenAndServe("localhost:8080", nil)
}
