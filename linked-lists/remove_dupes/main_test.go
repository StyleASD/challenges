package main

import (
	"testing"

	list "programming-problems/linked-lists"
)

func TestRemoveDupes_FirstAttempt(t *testing.T) {
	type args struct {
		linkedList *list.SinglyLinkedList
	}
	tests := []struct {
		name   string
		args   args
		length int
	}{
		{name: "Test Case 1", args: args{linkedList: list.NewSinglyLinkedList()}, length: 0},
		{name: "Test Case 2", args: args{linkedList: list.NewSinglyLinkedList(1, 1)}, length: 1},
		{name: "Test Case 3", args: args{linkedList: list.NewSinglyLinkedList(1, 2, 3, 4, 1, 5, 6)}, length: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RemoveDupes_FirstAttempt(tt.args.linkedList)
			if tt.args.linkedList.Length != tt.length {
				t.Errorf("RemoveDupes_FirstAttempt() = %v, want %v", tt.args.linkedList.Length, tt.length)
			}
		})
	}
}
