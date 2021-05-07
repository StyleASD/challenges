package main

import (
	"reflect"
	"testing"

	linked_lists "programming-problems/linked-lists"
)

func TestIntersection_Attempt1(t *testing.T) {
	type args struct {
		list1 linked_lists.SinglyLinkedList
		list2 linked_lists.SinglyLinkedList
	}

	f := func(name string) struct {
		name string
		args args
		want *linked_lists.SingleNode
	} {
		node := &linked_lists.SingleNode{
			Data: 3,
		}
		list1 := linked_lists.NewSinglyLinkedList(1, 3, 4, 5)
		list2 := linked_lists.NewSinglyLinkedList(1, 8)

		list1.AppendNode(node)
		list1.AppendData(3)
		list1.AppendData(4)
		list2.AppendNode(node)

		return struct {
			name string
			args args
			want *linked_lists.SingleNode
		}{
			name: name,
			args: args{
				list1: *list1,
				list2: *list2,
			},
			want: node,
		}
	}

	tests := []struct {
		name string
		args args
		want *linked_lists.SingleNode
	}{
		f("Test Case 1"),
		{
			name: "Test Case 2",
			args: args{
				list1: *linked_lists.NewSinglyLinkedList(1, 2, 3),
				list2: *linked_lists.NewSinglyLinkedList(1, 2, 3),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection_Attempt1(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection_Attempt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
