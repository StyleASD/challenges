package main

import (
	"testing"

	linked_lists "programming-problems/linked-lists"
)

func TestSumList_Attempt1(t *testing.T) {
	type args struct {
		list1 linked_lists.SinglyLinkedList
		list2 linked_lists.SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want linked_lists.SinglyLinkedList
	}{
		{
			name: "Example 1",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(7, 1, 6),
				list2: *linked_lists.NewSinglyLinkedList(5, 9, 2),
			},
			want: *linked_lists.NewSinglyLinkedList(2, 1, 9),
		},
		{
			name: "Example 2",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(1, 2, 3),
				list2: *linked_lists.NewSinglyLinkedList(3, 2, 1),
			},
			want: *linked_lists.NewSinglyLinkedList(4, 4, 4),
		},
		{
			name: "Example 3",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(0, 5, 5),
				list2: *linked_lists.NewSinglyLinkedList(0, 5, 6),
			},
			want: *linked_lists.NewSinglyLinkedList(0, 0, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SumListReverse_Attempt1(tt.args.list1, tt.args.list2).Head

			expect := tt.want.Head

			for expect != nil {
				if expect.Data != actual.Data {
					t.Errorf("SumListReverse_Attempt1() = %v, want %v", actual.Data, expect.Data)
				}
				expect = expect.Next
				actual = actual.Next
			}

		})
	}
}

func TestSumListForward_Attempt1(t *testing.T) {
	type args struct {
		list1 linked_lists.SinglyLinkedList
		list2 linked_lists.SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want linked_lists.SinglyLinkedList
	}{
		{
			name: "Example 1",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(6, 1, 7),
				list2: *linked_lists.NewSinglyLinkedList(2, 9, 5),
			},
			want: *linked_lists.NewSinglyLinkedList(9, 1, 2),
		},
		{
			name: "Example 2",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(1, 2, 3),
				list2: *linked_lists.NewSinglyLinkedList(3, 2, 1),
			},
			want: *linked_lists.NewSinglyLinkedList(4, 4, 4),
		},
		{
			name: "Example 3",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(5, 5, 0),
				list2: *linked_lists.NewSinglyLinkedList(6, 5, 0),
			},
			want: *linked_lists.NewSinglyLinkedList(1, 2, 0, 0,),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SumListForward_Attempt1(tt.args.list1, tt.args.list2).Head

			expect := tt.want.Head

			for expect != nil {
				if expect.Data != actual.Data {
					t.Errorf("SumListForward_Attempt1() = %v, want %v", actual.Data, expect.Data)
				}
				expect = expect.Next
				actual = actual.Next
			}
		})
	}
}
