package main

import "fmt"

type Stack[T comparable] struct {
	elements []T
}

func (stack *Stack[T]) Push(element T) {
	stack.elements = append(stack.elements, element)
}
func (stack *Stack[T]) Empty() bool {
	return len(stack.elements) == 0
}
func (stack *Stack[T]) Pop() T {
	if stack.Empty() {
		panic("stack is empty")
	}
	element := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return element
}

func (stack *Stack[T]) Peek() T {
	if stack.Empty() {
		panic("stack is empty")
	}
	return stack.elements[len(stack.elements)-1]
}
func (stack *Stack[T]) Size() int {
	return len(stack.elements)
}
func (stack *Stack[T]) Clear() {
	stack.elements = []T{}
}
func main() {
	stack := Stack[int]{}
	stack.Push(3)
	fmt.Println(stack.Peek())
	stack.Push(2)
	fmt.Println(stack.Peek())
	stack.Pop()
	fmt.Println(stack.Size())
	fmt.Println(stack.Peek())
}
