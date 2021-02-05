package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type PostHandler struct{}

func (h PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, readErr := ioutil.ReadAll(r.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		_, writeErr := io.WriteString(w, string(body))
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", new(PostHandler)))
}
