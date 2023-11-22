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

func TestParseRow(t *testing.T) {
	tests := []struct {
		name       string
		row        string
		wantCount  int
		wantValues []string
	}{
		{"simple", `
		<tr data-v-7bc433c8><td class="rounded-l-2 px-4 md:px-5 py-3 border-0 border-r-1 border-r-white border-solid" data-v-7bc433c8><!---->
										Nov. 18
									</td> <td class="px-4 md:px-5 py-3 border-0 border-r-1 border-r-white border-solid" data-v-7bc433c8>
										882
									</td> <td class="rounded-r-2 px-4 md:px-5 py-3 font-bold" data-v-7bc433c8>
											THINK
										</td></tr>
`,
			3,
			[]string{
				"2023-11-18",
				"882",
				"THINK",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cells := ParseRow(tt.row)
			assert.Equal(t, tt.wantCount, len(cells))
			assert.Equal(t, tt.wantValues, cells)
		})
	}
}

/*
 */
