package stack

import "testing"

func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	if minStack.GetMin() != -3 {
		t.Errorf("Expected GetMin to return -3, got %d", minStack.GetMin())
	}

	minStack.Pop()

	if minStack.Top() != 0 {
		t.Errorf("Expected Top to return 0, got %d", minStack.Top())
	}

	if minStack.GetMin() != -2 {
		t.Errorf("Expected GetMin to return -2, got %d", minStack.GetMin())
	}
}
