package main

import (
	"flag"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4040", "HTTP address")
	app := &Application{
		logInfo:  InfoLogger(),
		logError: ErrorLogger(),
	}

	flag.Parse()

	server := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: app.logError,
	}

	app.logInfo.Printf("Server started on port %v!\n", *addr)
	err := server.ListenAndServe()
	app.logError.Fatal(err)
}
