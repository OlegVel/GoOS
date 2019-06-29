package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("files"))
	http.Handle("/files/", http.StripPrefix("/files/", fileServer))
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}
