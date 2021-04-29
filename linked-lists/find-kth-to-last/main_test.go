package main

import (
	"testing"

	list "programming-problems/linked-lists"
)

func TestFindKthToLast_Attempt1(t *testing.T) {
	type args struct {
		k   int
		sll list.SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test Case 1", args: args{sll: *list.NewSinglyLinkedList(), k: 5}, want: 0},
		{name: "Test Case 2", args: args{sll: *list.NewSinglyLinkedList(2, 1), k: 2}, want: 2},
		{name: "Test Case 3", args: args{sll: *list.NewSinglyLinkedList(1, 2, 3, 4, 1, 5, 6), k: 5}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthToLast_Attempt1(tt.args.k, tt.args.sll); got != tt.want {
				t.Errorf("FindKthToLast_Attempt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
