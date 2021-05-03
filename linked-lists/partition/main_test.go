package main

import (
	"testing"

	linked_lists "programming-problems/linked-lists"
)

func TestPartition_Attempt1(t *testing.T) {
	type args struct {
		sll       *linked_lists.SinglyLinkedList
		partition int
	}
	tests := []struct {
		name     string
		args     args
		expected *linked_lists.SinglyLinkedList
	}{
		{
			name: "Example Case",
			args: args{
				sll:       linked_lists.NewSinglyLinkedList(3, 5, 8, 10, 5, 2, 1),
				partition: 5,
			},
			expected: linked_lists.NewSinglyLinkedList(1, 2, 3, 5, 8, 10, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Partition_Attempt1(tt.args.sll, tt.args.partition)

			expect := tt.expected.Head
			actual := tt.args.sll.Head

			for expect != nil {
				if expect.Data != actual.Data {
					t.Errorf("Partition_Attempt1() = %v, want %v", actual.Data, expect.Data)
				}
				expect = expect.Next
				actual = actual.Next
			}

		})
	}
}
