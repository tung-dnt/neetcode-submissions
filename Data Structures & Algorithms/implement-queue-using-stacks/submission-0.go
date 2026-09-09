type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}}
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	} 
	peek := this.Peek()
	this.list = this.list[1:]
	return peek
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.list) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Peek();
 * param4 := obj.Empty();
 */
