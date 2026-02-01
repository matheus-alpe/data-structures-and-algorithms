package bitwiseoperators

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "Example 1",
			num:      11, // Binary: 1011
			expected: 3,
		},
		{
			name:     "Example 2",
			num:      128, // Binary: 10000000
			expected: 1,
		},
		{
			name:     "Example 3",
			num:      4294967293, // Binary: 11111111111111111111111111111101
			expected: 31,
		},
		{
			name:     "All bits zero",
			num:      0, // Binary: 0
			expected: 0,
		},
		{
			name:     "All bits one (8 bits)",
			num:      255, // Binary: 11111111
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hammingWeight(tt.num)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
