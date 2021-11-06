package main

import (
	"log"
	"net/http"
)

func main() {
	fileSrv := http.FileServer(http.Dir("."))
	log.Fatalln(http.ListenAndServe(":8080", fileSrv))
}
