package main

import (
	"testing"

	linked_lists "programming-problems/linked-lists"
)

func TestLoopDetection_Attempt1(t *testing.T) {
	type args struct {
		list1 linked_lists.SinglyLinkedList
	}
	tests := []struct {
		name string
		args args
		want *linked_lists.SingleNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := LoopDetection_Attempt1(tt.args.list1)

			if tt.want != node {
				t.Errorf("Partition_Attempt1() = %v, want %v", node, tt.want)
			}
		})
	}
}
