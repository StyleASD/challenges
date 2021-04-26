package main

import (
	"testing"
)

func TestIsRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test Case 1", args{"waterbottle", "erbottlewat"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
