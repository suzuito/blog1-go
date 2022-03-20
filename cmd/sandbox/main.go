package main

import (
	"github.com/piprate/json-gold/ld"
)

func main() {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	doc := map[string]interface{}{
		"headline": "hog",
	}

	context := map[string]interface{}{
		"@context": "https://schema.org",
		"@type":    "Article",
	}

	compactedDoc, err := proc.Compact(doc, context, options)
	if err != nil {
		panic(err)
	}

	ld.PrintDocument("JSON-LD compation succeeded", compactedDoc)
}
