package modeler

import (
	"testing"

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
		})
	}
}
