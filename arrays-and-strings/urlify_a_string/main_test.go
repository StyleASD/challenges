package main

import (
	"reflect"
	"testing"
)

func TestUrlify(t *testing.T) {
	type args struct {
		strArray   string
		trueLength int
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"Example", args{"Mr John Smith    ", 13}, []rune("Mr%20John%20Smith")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Urlify([]rune(tt.args.strArray), tt.args.trueLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Urlify() = %v, want %v", got, tt.want)
			}
		})
	}
}