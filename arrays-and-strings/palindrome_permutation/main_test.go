package main

import (
	"testing"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test Input",
			args{"Tact Coa"},
			true,
		},
		{
			"Longer test",
			args{"never odd or even"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutationOfPalindrome(tt.args.str); got != tt.want {
				t.Errorf("IsPermutationOfPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
