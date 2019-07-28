package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("files"))
	http.Handle("/files/", http.StripPrefix("/files/", fileServer))
	http.Handle("/", http.FileServer(http.Dir("front")))
	http.HandleFunc("/api", handleApi())

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}

type testStruct struct {
	Test string
}

func handleApi() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var t testStruct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}
		log.Println(t.Test)

		fmt.Println("Hello world")
	}
}
