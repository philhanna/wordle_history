package dbcreator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToYYYYMMDD(t *testing.T) {
	type args struct {
		yymmdd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid date", args{"230621"}, "2023-06-21"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToYYYYMMDD(tt.args.yymmdd); got != tt.want {
				t.Errorf("ToYYYYMMDD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetScrapes(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		expected int
	}{
		{"from test data", "testdata/answers.html", 886},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := os.ReadFile(tt.filename)
			assert.Nil(t, err)
			scrapes := GetScrapes(string(body))
			assert.Equal(t, tt.expected, len(scrapes))
		})
	}
}
