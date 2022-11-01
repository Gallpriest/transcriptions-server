package main

import "net/http"

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./client/static/"))

	mux.HandleFunc("/history", app.historyRoute)
	mux.HandleFunc("/search", app.getSearchResult)
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}
