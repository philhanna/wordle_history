package main

import (
// "time"
	"regexp"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

/*
input := "2017-08-31"
layout := "2006-01-02"
t, _ := time.Parse(layout, input)
fmt.Println(t.Format("02-Jan-2006")) // Output: 31-Aug-2017
*/
/*
type Scraper struct {
	date   string // Date word was used, format YYYY-MM-DD
	puzzle int    // Puzzle number
	word   string // The word
}
*/

func GetRows(body string) []string {
	re := regexp.MustCompile(`(?s)(<tr.*?</tr>)`)
	return re.FindAllString(body, -1)
}