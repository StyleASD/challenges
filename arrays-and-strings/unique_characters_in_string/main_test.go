package main

import (
	"testing"
)

func TestStringHasUniqueCharacters(t *testing.T) {

	/**
	- aadvark = false
	- and     = true
	- Zoe     = true
	- ZEAzPlp = true
	 */
	type args struct {
		str string
	}
	tests := []struct {
		name string
		str string
		want bool
	}{
		{"aardvark should be false", "aardvark", false},
		{"and should be true", "and", true},
		{"Zoe should be true", "Zoe", true},
		{"ZEAzPlp should be true", "ZEAzPlp", true},
		{"/abcfts/ should be false", "/abcfts/", false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringHasUniqueCharacters_AttemptOne(tt.str); got != tt.want {
				t.Errorf("StringHasUniqueCharacters_AttemptOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
