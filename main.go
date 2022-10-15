package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func main() {
	http.HandleFunc("/", HelloHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
