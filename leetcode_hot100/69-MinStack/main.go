package main

func main() {

}

type MinStack struct {
	Stack    []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{Stack: make([]int, 0), MinStack: []int{(1 << 63) - 1}}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	this.MinStack = append(this.MinStack, min(this.MinStack[len(this.MinStack)-1], val))
}

func (this *MinStack) Pop() {
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
	this.Stack = this.Stack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}
