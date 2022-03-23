package main

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

//Implement the MinStack class:

//MinStack() initializes the stack object.
//void push(int val) pushes the element val onto the stack.
//void pop() removes the element on the top of the stack.
//int top() gets the top element of the stack.
//int getMin() retrieves the minimum element in the stack.

type MinStack struct {
	myStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ms.myStack = append(ms.myStack, val)
}

func (ms *MinStack) Pop() {
	ms.myStack = ms.myStack[:len(ms.myStack)-1]
}

func (ms *MinStack) Top() int {
	result := ms.myStack[len(ms.myStack)-1]
	return result
}

func (ms *MinStack) GetMin() int {
	min := ms.myStack[0]
	for i := 1; i < len(ms.myStack); i++ {
		if ms.myStack[i] < min {
			min = ms.myStack[i]
		}
	}
	return min
}
