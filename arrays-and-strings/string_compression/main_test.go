package main

import (
	"testing"
)

func TestCompressString_Attempt1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"aabcccccaaa"}, "a2b1c5a3"},
		{"Test Case 1", args{"abc"}, "abc"},
		{"Test Case 2", args{"abbcc"}, "abbcc"},
		{"Test Case 3", args{"abbbcc"}, "abbbcc"},
		{"Test Case 3", args{"aabbbbcc"}, "a2b4c2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString_Attempt1(tt.args.str); got != tt.want {
				t.Errorf("CompressString_Attempt1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompressString_Attempt1(b *testing.B) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"aabcccccaaa"}, "a2b1c5a3"},
		{"Test Case 1", args{"abc"}, "abc"},
		{"Test Case 2", args{"abbcc"}, "abbcc"},
		{"Test Case 3", args{"abbbcc"}, "abbbcc"},
		{"Test Case 3", args{"aabbbbcc"}, "a2b4c2"},
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			CompressString_Attempt1(tt.args.str)
		}
	}
}
