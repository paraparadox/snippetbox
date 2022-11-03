package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// don't forget about 2 types of URL patterns:
	// fixed paths and subtree paths
	// mux.HandleFunc("/really/serious", someFunc)
	// mux.HandleFunc("/really/serious/", otherFunc)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
