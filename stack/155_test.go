package stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	min := minStack.GetMin()
	if min != -2 {
		t.Errorf("GetMin() = %v; want -3", min)
	}
	top := minStack.Top()
	if top != -1 {
		t.Errorf("GetMin() = %v; want 0", top)
	}
	minStack.Pop()
	min = minStack.GetMin()
	if min != -2 {
		t.Errorf("GetMin() = %v; want -2", min)
	}
}
