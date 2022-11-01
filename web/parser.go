package main

import (
	"reflect"

	"golang.org/x/net/html"
)

func ParseTranscription(n *html.Node, nodes *ParserResult) {
	if n.Type == html.ElementNode && n.Data == "span" {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "hw dhw" {
				nodes.Name = *n
				break
			} else if a.Key == "class" && a.Val == "ipa dipa lpr-2 lpl-1" {
				nodes.Transcription = *n
				break
			} else {
				continue
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !reflect.ValueOf(nodes.Name).IsZero() && !reflect.ValueOf(nodes.Transcription).IsZero() {
			break
		}
		ParseTranscription(c, nodes)
	}
}
