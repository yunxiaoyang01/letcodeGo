package main

type CQueue struct {
	inStack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

func (this *CQueue) DeleteHead() int {
	l := len(this.inStack)
	if l == 0 {
		return -1
	}
	val := this.inStack[0]
	this.inStack = this.inStack[1:]
	return val
}

func main() {
	queue := Constructor()
	queue.DeleteHead()
	queue.AppendTail(5)
	queue.AppendTail(2)
	queue.DeleteHead()
	queue.DeleteHead()
}
