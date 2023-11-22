package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRows(t *testing.T) {
	tests := []struct {
		name string
		body string
		want []string
	}{
		{
			"simple",
			`Something
more something <TR x="1">body1</tr><tr>body2</tr>`,
			[]string{
				`<TR x="1">body1</tr>`,
				`<tr>body2</tr>`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GetRows(tt.body)
			expected := tt.want
			assert.Equal(t, expected, actual)
		})
	}
}
