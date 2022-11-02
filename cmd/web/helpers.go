package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError writes an error message and stack trace to the errorLog and sends a generic 500 response
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and description to the user; used when there's a problem with the request that the user sent
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is simply a wrapper around clientError which sends 404 response; implemented for consistency
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
