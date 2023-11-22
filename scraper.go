package main

import (
	// "time"
	"regexp"
	"strings"
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
	re := regexp.MustCompile(`(?si)(<tr.*?</tr>)`)
	return re.FindAllString(body, -1)
}

func ParseRow(row string) []string {
	re := regexp.MustCompile(`(?s)<td.*?>(.*?)</td>`)
	m := re.FindAllStringSubmatch(row, -1)
	if m == nil {
		return nil
	}
	output := make([]string, 0)
	for _, cell := range m {
		td := cell[1]
		td = strings.ReplaceAll(td, "<!---->", "")
		trimmed := strings.Trim(td, "\n\t ")
		output = append(output, trimmed)
	}
	return output
}