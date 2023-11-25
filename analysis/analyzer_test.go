package analysis

import (
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestLoadWords(t *testing.T) {
	tests := []struct {
		name          string
		filename      string
		wantCount     int
		wantFirstWord string
		wantLastWord  string
	}{
		{"Check count and first word", "testdata/history.db", 889, "CIGAR", "THROW"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			words, err := LoadWords(tt.filename)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantCount, len(words))
			assert.Equal(t, tt.wantFirstWord, words[0])
			assert.Equal(t, tt.wantLastWord, words[len(words)-1])
			for i := 0; i < 10; i++ {
				fmt.Printf("%d: %s\n", i, words[i])
			}
		})
	}
}

func TestLetterFrequency(t *testing.T) {
	tests := []struct {
		name string
		filename string
		want Statistic
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
