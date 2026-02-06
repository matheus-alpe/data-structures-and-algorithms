package stack

type MinStack struct {
	top *NodeMin
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	node := &NodeMin{value: val, currentMin: val}
	if this.top != nil {
		node.next = this.top
		node.currentMin = min(val, this.top.currentMin)
	}
	this.top = node
}

func (this *MinStack) Pop() {
	if this.top != nil {
		this.top = this.top.next
	}
}

func (this *MinStack) Top() int {
	return this.top.value
}

func (this *MinStack) GetMin() int {
	return this.top.currentMin
}

type NodeMin struct {
	value      int
	currentMin int
	next       *NodeMin
}
