package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func EndHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "End World!!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("template No Such")
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
