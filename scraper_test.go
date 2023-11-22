package main

import "testing"

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
