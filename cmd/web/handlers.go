package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//initialoize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must always be the first file in the slice.

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	//use the template.ParseFiles() function to read the template file into a
	// template set. Notice that we can pass the slice of file paths as a variadic
	// parameter? This is a handy feature of Go that allows us to pass a variable
	// number of arguments of the same type to a function. If there's an error,
	// we use the log.Fatal() function to log the error message and exit.

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	//we then use the Execute() method on the template set to write the template
	//content as the respose body. The last parameter to Execute() represents
	//any dynamic data that we want to pass in, which for now will be nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./ui/static/file.zip")
	w.Header().Set("Content-Disposition", "attachment; filename=downloaded.txt")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))
	w.Write([]byte("This is the content of the file"))
}
