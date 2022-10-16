package main

import (
	"fmt"
	"log"
	"net/http"
)

func EndHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "End World!!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/end", EndHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
