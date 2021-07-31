package main

import (
	"fmt"
	"log"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sup, nibba, you are at %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

