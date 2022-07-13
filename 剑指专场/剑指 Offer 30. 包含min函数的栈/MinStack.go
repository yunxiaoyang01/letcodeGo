package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack        []int
	min, lastMin int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min:     math.MaxInt,
		lastMin: math.MaxInt,
	}
}

func (this *MinStack) Push(x int) {
	if this.min > x {
		this.lastMin, this.min = this.min, x
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if val == this.min {
		this.min = this.lastMin
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	if len(this.stack) == 1 {
		return this.stack[0]
	}
	return this.min
}

func main() {
	minStack := Constructor()
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	fmt.Printf("%d", minStack.Min())
	minStack.Pop()
	fmt.Printf("%d", minStack.Min())
}
