package dynamicprogramming

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{2, 1, 1, 2}, 4},
		{[]int{}, 0},
		{[]int{5}, 5},
		{[]int{1, 2}, 2},
	}

	for _, test := range tests {
		result := rob(test.nums)
		if result != test.expected {
			t.Errorf("rob(%v) = %d; expected %d", test.nums, result, test.expected)
		}
	}
}
