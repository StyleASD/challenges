package main

import (
	"reflect"
	"testing"
)

func TestZeroMatrix_Attempt1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test case 1",
			args{
				[][]int{
					{1, 2, 3, 4, 5},
					{1, 0, 2, 4, 5},
					{1, 2, 3, 4, 5},
					{1, 2, 3, 4, 5},
				},
			},
			[][]int{
				{1, 0, 3, 4, 5},
				{0, 0, 0, 0, 0},
				{1, 0, 3, 4, 5},
				{1, 0, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroMatrix_Attempt1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroMatrix_Attempt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
