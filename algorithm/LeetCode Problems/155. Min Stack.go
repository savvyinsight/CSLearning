package main

// 155. Min Stack
// https://leetcode.com/problems/min-stack/

/*
1. Stack
*/

type MinStack struct {
	st  []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.st = append(this.st, val)
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.st) == 0 {
		return
	}
	val := this.st[len(this.st)-1]
	this.st = this.st[:len(this.st)-1]
	if val == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.st) == 0 {
		return 0
	}
	return this.st[len(this.st)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	println(obj.GetMin()) // return -3
	obj.Pop()
	println(obj.Top())    // return 0
	println(obj.GetMin()) // return -2
}
