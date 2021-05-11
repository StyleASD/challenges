package main

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
	count         [...]int
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
		count:         make([...]int, 3),
	}
}

func (t ThreeInOne_Attempt1) Pop(stack StackNo) interface{} {
	if !stack.isValid() {
		return nil
	}

	return nil
}
