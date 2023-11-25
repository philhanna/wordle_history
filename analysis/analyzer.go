package analysis

import (
	"database/sql"
	"sort"

	_ "github.com/mattn/go-sqlite3"
)

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------
type Statistic struct {
	Name  string
	Value string
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Analyze reads a list of words and computes statistics from it.
func Analyze(words []string) []Statistic {
	return nil
}

// LetterFrequency scans the list of words and returns a string
// containing the letters they contain in descending order of frequency.
func LetterFrequency(words []string) Statistic {
	freq := make(map[byte]int)
	for _, word := range words {
		for i := 0; i < WORDLEN; i++ {
			ch := word[i]
			freq[ch]++
		}
	}
	// Sort the map descending by frequency count
	type LetterCount struct {
		letter byte
		count  int
	}
	letterCounts := make([]LetterCount, 0)
	for letter, count := range freq {
		letterCounts = append(letterCounts, LetterCount{letter, count})
	}
	sort.Slice(letterCounts, func(i, j int) bool {
		return letterCounts[i].count > letterCounts[j].count
	})

	// Assemble the letters into a string
	letters := make([]byte, len(letterCounts))
	for i := 0; i < len(letterCounts); i++ {
		letters[i] = letterCounts[i].letter
	}
	lettersString := string(letters)

	stat := Statistic{
		Name:  "Letter frequency",
		Value: lettersString,
	}
	return stat
}

// LoadWords get the list of words from the history database
func LoadWords(filename string) ([]string, error) {

	// Connect to the database
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Issue a query for the words in the history table
	rows, err := db.Query(`SELECT word FROM history ORDER BY date`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read from the query result into a slice of strings
	words := make([]string, 0)
	for rows.Next() {
		var word string
		if err := rows.Scan(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
	}

	return words, nil
}
