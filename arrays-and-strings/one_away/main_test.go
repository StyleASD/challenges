package main

import (
	"testing"
)

func TestOneAway(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"pale", "ple"}, true},
		{"Example 2", args{"pales", "pale"}, true},
		{"Example 3", args{"pale", "bale"}, true},
		{"Example 4", args{"pale", "bake"}, false},
		{"Test Case 1", args{"pal", "paler"}, false},
		{"Test Case 2", args{"paler", "pal"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneAway(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
