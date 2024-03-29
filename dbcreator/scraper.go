package dbcreator

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

type MonthData struct {
	Month   string       `json:"month"`
	Answers []AnswerData `json:"answers"`
}

type AnswerData struct {
	Date   string `json:"date"`
	Index  string `json:"index"`
	Answer string `json:"answer"`
}

// GetScrapes extracts the answer data from the specified string.
// This data is in JSON-like format following the string "pastData:".
// JSON-like because it is not quite valid JSON for two reasons:
//
// 1. The property names are not quoted.
//
// 2. The index field is almost always numeric, but there are several
// places where it is unquoted alphabetic data.
//
// So after the JSON string is extracted, I pre-process the
// property names and index values to make them quoted.
func GetScrapes(body string) []AnswerData {

	log.Printf("Parsing HTML for answer data...\n")

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
	body = "[" + body[p:q]

	// Quote the value of index:
	re := regexp.MustCompile(`index:(.*?),`)
	body = re.ReplaceAllString(body, `index:"$1",`)

	// Quote property names
	re = regexp.MustCompile(`(\w+):`)
	body = re.ReplaceAllString(body, `"$1":`)

	// Now unmarshal the JSON structure into a slice of MonthData
	var data []MonthData
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatal(err)
	}

	// Unwrap the answer data from each month
	answers := make([]AnswerData, 0)
	for _, monthStruct := range data {
		answers = append(answers, monthStruct.Answers...)
	}

	// Convert the date string YYMMDD to YYYY-MM-DD
	for i, answer := range answers {
		answer.Date = ToYYYYMMDD(answer.Date)
		answers[i] = answer
	}
	return answers
}

// ToYYYYMMDD reformats the date string from yymmdd to YYYY-MM-DD
func ToYYYYMMDD(yymmdd string) string {
	yy, _ := strconv.Atoi(yymmdd[:2])
	mm, _ := strconv.Atoi(yymmdd[2:4])
	dd, _ := strconv.Atoi(yymmdd[4:])
	yyyymmdd := fmt.Sprintf("20%02d-%02d-%02d", yy, mm, dd)
	return yyyymmdd
}
