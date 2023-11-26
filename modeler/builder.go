package modeler

import (
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
type LetterCount struct {
	letter byte
	count  int
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Analyze reads a list of words and computes statistics from it.
func Analyze(words []string) []Statistic {
	return nil
}

// GetLetterFrequencies scans the list of words and returns a slice of
// LetterCount structure that describe equal letter found in the list
// and the number of times it occurs
func GetLetterFrequencies(words []string) []LetterCount {

	// Compute the frequencies
	freq := make(map[byte]int)
	for _, word := range words {
		for i := 0; i < WORDLEN; i++ {
			ch := word[i]
			freq[ch]++
		}
	}

	// Make a list of letter count structures
	letterCounts := make([]LetterCount, 0)
	for letter, count := range freq {
		letterCounts = append(letterCounts, LetterCount{letter, count})
	}
	
	// Sort the list descending by frequency count, descending
	sort.Slice(letterCounts, func(i, j int) bool {
		return letterCounts[i].count > letterCounts[j].count
	})

	return letterCounts
}

// LetterFrequency scans the list of words and returns a string
// containing the letters they contain in descending order of frequency.
func LetterFrequency(words []string) Statistic {

	letterCounts := GetLetterFrequencies(words)

	// Assemble the letters into a string
	letters := make([]byte, len(letterCounts))
	for i := 0; i < len(letterCounts); i++ {
		letters[i] = letterCounts[i].letter
	}
	lettersString := string(letters)

	// Return the computed statistic
	stat := Statistic{
		Name:  "Letter frequency",
		Value: lettersString,
	}
	return stat
}
