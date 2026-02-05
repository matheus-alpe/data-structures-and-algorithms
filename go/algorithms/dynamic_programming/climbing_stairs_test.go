package dynamicprogramming

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{10, 89},
	}

	for _, test := range tests {
		result := climbStairs(test.n)
		if result != test.expected {
			t.Errorf("climbStairs(%d) = %d; expected %d", test.n, result, test.expected)
		}
	}
}
