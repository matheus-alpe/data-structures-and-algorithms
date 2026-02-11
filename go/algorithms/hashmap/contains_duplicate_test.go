package hashmap

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{nums: []int{1, 2, 3, 4}, expected: false},
		{nums: []int{1, 2, 3, 1}, expected: true},
		{nums: []int{}, expected: false},
		{nums: []int{5, 5, 5, 5}, expected: true},
		{nums: []int{-1, -2, -3, -1}, expected: true},
	}

	for _, test := range tests {
		result := containsDuplicate(test.nums)
		if result != test.expected {
			t.Errorf("containsDuplicate(%v) = %v; want %v", test.nums, result, test.expected)
		}
	}
}
