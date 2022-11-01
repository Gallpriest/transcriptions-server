package main

import "golang.org/x/net/html"

type Data struct {
	Search string
	Links  [3]string
}

type SearchResult struct {
	Value string
}

type ParserResult struct {
	Name          html.Node
	Transcription html.Node
}

type JSONSearchResult struct {
	Name          string
	Transcription string
}
