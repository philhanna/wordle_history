package main

import (
	"os"
	"regexp"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type Scrape struct {
	date   string // Date word was used, format YYYY-MM-DD
	puzzle int    // Puzzle number
	word   string // The word
}

func GetScrapes(body string) []Scrape {
	
	const startTag = "pastData:["

	// Find the start of the past data
	p := strings.Index(body, startTag)
	if p == -1 {
		return nil
	}
	p += len(startTag)
	q := p

	// Find the end of the past data by finding the last right bracket
	// at this level
	for level := 1; level > 0; q++ {
		switch body[q] {
		case '[':
			level++
		case ']':
			level--
		default:
		}
	}

	// q now points to the last closing bracket at this level
	subBody := "[" + body[p:q]

	re := regexp.MustCompile(`(\w+):`)
	a := re.ReplaceAllString(subBody, `"$1"`)

	os.WriteFile("/tmp/a.json", []byte(a), 0644)


	scrapes := make([]Scrape, 0)
	return scrapes
}
