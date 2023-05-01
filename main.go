package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))

}
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//this should be before any call to other methods
		w.Header().Set("Allow", http.MethodPost)
		// to send a non-200 status code,
		// you must call w.WriteHeader() before any call to w.Write().

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	w.Header()["Date"] = nil
	w.Write([]byte("Create a new snippet..."))
}

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
