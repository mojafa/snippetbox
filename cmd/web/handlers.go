package main

import (
	"errors"
	"fmt"
	"mojafa/snippetbox/pkg/models"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) //use notFound helper
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) //use serverError helper
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) //use serverError helper
	}
}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	// Use the SnippetModel object's Get method to retrieve the data for a // specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%v", s)
	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//this should be before any call to other methods
		w.Header().Set("Allow", http.MethodPost)
		// to send a non-200 status code,
		// you must call w.WriteHeader() before any call to w.Write().

		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Header()["X-XSS-Protection"] = []string{"1; mode=block"}
	// w.Header()["Date"] = nil
	// w.Write([]byte("Create a new snippet..."))
	// Create some variables holding dummy data. We'll remove these later on
	// 	during the build.
	title := "1984"
	content := "George Orwells best big brother"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/file.zip")
}
