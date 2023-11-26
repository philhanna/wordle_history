package modeler

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestLetterFrequency(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     Statistic
	}{
		{"Full list", "testdata/history.db", Statistic{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words, err := LoadWords(tt.filename)
			assert.Nil(t, err)
			actual := LetterFrequency(words)
			fmt.Printf("%s\n", actual.Value)
		})
	}
}

func TestGetLetterFrequencies(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []LetterCount
	}{
		{"two words", []string{"PIZZA", "QUICK"}, []LetterCount{
			{'P', 1},
			{'I', 2},
			{'Z', 2},
			{'A', 1},
			{'Q', 1},
			{'U', 1},
			{'C', 1},
			{'K', 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := GetLetterFrequencies(tt.words)
			assert.ElementsMatch(t, tt.want, have)
		})
	}
}
