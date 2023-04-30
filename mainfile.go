package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))

}

// Add a showSnippet handler function.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	//Extract th value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	//use the fmt.Printf() function to interpolate the id value with our
	// response and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	//w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status code an
	// the w.Write() method to write a "Method Not Allowed" response body. We
	// then return from the function so that the subsequent code is not execute

	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header to the
		// response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		//w.Header().Set("Allow", "POST")
		w.Header().Set("Allow", "POST")
		//suppresisng the date header content
		//w.Header()["Date"] = nil
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))
		// Use the http.Error() function to send a 405 status code and "Method Allowed" string as the response body.

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern. // Register the two new handler functions and corresponding URL patterns wi
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server. We pas
	// two parameters: the TCP network address to listen on (in this case ":400
	// and the servemux we just created. If http.ListenAndServe() returns an er
	// we use the log.Fatal() function to log the error message and exit.
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
