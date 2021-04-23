package main

import (
	"reflect"
	"testing"
)

func TestRotateMartix_Attempt1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name        string
		args        args
		wantRotated [][]int
	}{
		{"Test Case 1", args{
			[][]int{
				{1, 2},
				{3, 4},
			},
		},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
		{"Test Case 2", args{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{"Test Case 3", args{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
		},
			[][]int{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRotated := RotateMatrix_Attempt1(tt.args.matrix); !reflect.DeepEqual(gotRotated, tt.wantRotated) {
				t.Errorf("RotateMatrix_Attempt1() = %v, want %v", gotRotated, tt.wantRotated)
			}
		})
	}
}
