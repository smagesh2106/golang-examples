package main

import (
	"errors"
	"fmt"
)

//Stack definition of stack object
type Stack struct {
	top        int
	capacity   int
	stackArray []interface{}
}

//init intialises a new stack object and returns a pointer for the same
func (stack *Stack) init(capacity int) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.stackArray = make([]interface{}, capacity)

	return stack
}

//NewStack creates a new stack object
func NewStack(capacity int) *Stack {
	return new(Stack).init(capacity)
}

//IsFull checks whether the stack has reached it's capacity or not
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

//IsEmpty checks whether the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

//Size returns the current size of the Stack
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

//Push inserts the data in the top of the stack
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack overflow")
	}
	stack.top++
	stack.stackArray[stack.top] = data
	return nil
}

//Pop removes the data from the top of the stack
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	stack.top--
	return temp, nil
}

//Peek returns the data without removing from the top of stack
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	return temp, nil
}

//Clear removes all the elements in the stack
func (stack *Stack) Clear() {
	stack.stackArray = nil
	stack.top = -1
}

//Print prints all the stack element
func (stack *Stack) Print() {
	fmt.Print("[")
	for i := 0; i <= stack.top; i++ {
		fmt.Print(stack.stackArray[i])
		if i != stack.top {
			fmt.Print(",")
		}
	}
	fmt.Print("]\n")
}

func permutate2(a []rune, l int, callStack *Stack) {

	if l == len(a) {
		fmt.Println(string(a))
		callStack.Pop()
		callStack.Print()

		return
	}
	for i := l; i < len(a); i++ {
		callStack.Push(fmt.Sprintf("----------->>  a = %s, i = %d,  l = %d\n", string(a), i, l))
		a[i], a[l] = a[l], a[i]
		permutate2(a, l+1, callStack)
		a[i], a[l] = a[l], a[i]
	}
	//callStack.Pop()
	//callStack.Print()

}

var mystack *Stack

func main() {
	mystack = NewStack(25)
	//a := []rune("abc")
	permutate2([]rune("abc"), 0, mystack)
}
