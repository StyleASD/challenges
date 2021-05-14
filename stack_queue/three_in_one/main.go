package main

import (
	"fmt"
)

func main() {

}

type StackNo int

func (s StackNo) isValid() bool {
	return s > 0 && s < 4
}

/*
	Attempt 1:

	Assuming that we dont have to resize the stacks in the array we can split the array into three sections for the stack.
	If we need it to grow this would be much more complex.

	Answer:
	This is mostly correct apart apart form i haven't implemented yet. There is quite a few functions in the implementation.
	There is a much more extensive version that makes the arrays resizable. But this is very complex.
*/
type ThreeInOne_Attempt1 struct {
	stack         [...]int
	sizes         [...]int
	stackCapacity int
}

/*
	To achieve the above I am multiplying the input stack size by three. We can then use the size of each stack to as well
	as the stack number we want to access.
*/
func NewThreeInOne_Attempt1(stackCapacity int) *ThreeInOne_Attempt1 {
	return &ThreeInOne_Attempt1{
		stackCapacity: stackCapacity,
		stack:         make([...]int, stackCapacity*3),
		sizes:         make([...]int, 3),
	}
}

func (t *ThreeInOne_Attempt1) Push(data int, stack StackNo) error {
	if isFull, err := t.isFull(stack); isFull || err != nil {
		return err
	}

	// Increment stack pointer
	t.sizes[stack]++

	// Update the top value
	if pos, err := t.getTopIndex(stack); err == nil {
		t.stack[pos] = data
	}

	return nil
}

func (t *ThreeInOne_Attempt1) Pop(stack StackNo) (int, error) {
	if empty, err := t.IsEmpty(stack); empty || err != nil {
		return 0, err
	}

	pos, _ := t.getTopIndex(stack)

	data := t.stack[pos]
	t.stack[pos] = 0 // Clear the stack
	t.sizes[stack]--

	return data, nil
}

func (t *ThreeInOne_Attempt1) Peek(stack StackNo) (int, error) {
	if empty, err := t.IsEmpty(stack); empty || err != nil {
		return 0, err
	}

	pos, _ := t.getTopIndex(stack)
	return t.stack[pos], nil
}

func (t *ThreeInOne_Attempt1) IsEmpty(stack StackNo) (bool, error) {
	if err := validStack(stack); err != nil {
		return true, err
	}

	return t.sizes[stack] == 0, nil
}

func (t *ThreeInOne_Attempt1) isFull(stack StackNo) (bool, error) {
	if err := validStack(stack); err != nil {
		return true, err
	}

	return t.sizes[stack] == t.stackCapacity, nil
}

/**
 */
func (t *ThreeInOne_Attempt1) getTopIndex(stack StackNo) (int, error) {
	if err := validStack(stack); err != nil {
		return -1, err
	}

	offset := t.stackCapacity * int(stack)
	size := t.sizes[stack]

	return offset + size - 1, nil
}

func validStack(stack StackNo) error {
	if !stack.isValid() {
		return fmt.Errorf("Value %i is not a valid stack", stack)
	}

	return nil
}
