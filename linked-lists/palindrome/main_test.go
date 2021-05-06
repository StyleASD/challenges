package main

import (
	"testing"

	linked_lists "programming-problems/linked-lists"
)

func TestPalindrome_Attempt1(t *testing.T) {
	type args struct {
		sll linked_lists.SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				sll: *linked_lists.NewSinglyLinkedList(1, 2, 3, 2, 1),
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				sll: *linked_lists.NewSinglyLinkedList(1, 2, 3, 3, 2, 1),
			},
			want: true,
		},
		{
			name: "Test Case 3",
			args: args{
				sll: *linked_lists.NewSinglyLinkedList(1, 2, 3, 4, 2, 1),
			},
			want: false,
		},
		{
			name: "Test Case 4",
			args: args{
				sll: *linked_lists.NewSinglyLinkedList(1, 2),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Palindrome_Attempt1(tt.args.sll); got != tt.want {
				t.Errorf("Palindrome_Attempt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
