package main

import (
	"log"
	"net/http"
)

func main() {
	//locally scoped servemux.
	//This is a good practice to avoid polluting the global namespace.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting on server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
