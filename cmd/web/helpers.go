package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// the serverError() helper we use the debug.Stack() function to get a stack trace for the current goroutine and append it to the log messag
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// the clientError() helper we use the http.StatusText() function to automatically generate a human-friendly text representation of a given HTTP status code
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
