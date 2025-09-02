package stack

type StackItem struct {
	value int
	min   int
}

type MinStack struct {
	data []StackItem
}

func Constructor() MinStack {
	return MinStack{
		data: make([]StackItem, 0),
	}
}

func (this *MinStack) Push(val int) {
	valItem := StackItem{
		value: val,
		min:   val,
	}
	if len(this.data) == 0 {
		this.data = append(this.data, valItem)
		return
	}
	minNum := this.data[len(this.data)-1].min
	if val > minNum {
		valItem.min = minNum
	}
	this.data = append(this.data, valItem)
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	top := this.data[len(this.data)-1]
	return top.value
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1].min
}
