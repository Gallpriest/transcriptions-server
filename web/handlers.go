package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/html"
)

func (app *Application) historyRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("History route"))
}

func (app *Application) getSearchResult(w http.ResponseWriter, r *http.Request) {
	// EnableCors(&w)

	if r.Method != "POST" {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	result := SearchResult{}
	err := json.NewDecoder(r.Body).Decode(&result)

	if err != nil {
		app.serverError(w, err)
		app.logError.Fatal(err)
		return
	}

	nodes := ParserResult{}
	url := "https://dictionary.cambridge.org/dictionary/english/" + result.Value

	res, err := http.Get(url)

	if err != nil {
		app.serverError(w, err)
		app.logError.Fatal(err)
		return
	}

	data, _ := html.Parse(res.Body)
	ParseTranscription(data, &nodes)

	jsonStruct := &JSONSearchResult{
		Name:          string(nodes.Name.FirstChild.Data),
		Transcription: string(nodes.Transcription.FirstChild.Data),
	}
	response, _ := json.Marshal(jsonStruct)

	w.Write(response)
}
