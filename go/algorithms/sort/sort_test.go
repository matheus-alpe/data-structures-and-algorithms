package sort

import (
	"reflect"
	"testing"
)

// Common test cases for all sorting algorithms
var sortTestCases = []struct {
	name     string
	input    []int
	expected []int
}{
	{
		name:     "Unsorted array",
		input:    []int{64, 34, 25, 12, 22, 11, 90},
		expected: []int{11, 12, 22, 25, 34, 64, 90},
	},
	{
		name:     "Already sorted",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Empty array",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element",
		input:    []int{42},
		expected: []int{42},
	},
	{
		name:     "Duplicates",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
		expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
	},
	{
		name:     "All same elements",
		input:    []int{5, 5, 5, 5, 5},
		expected: []int{5, 5, 5, 5, 5},
	},
	{
		name:     "Negative numbers",
		input:    []int{-3, -1, -4, -1, -5},
		expected: []int{-5, -4, -3, -1, -1},
	},
	{
		name:     "Mixed positive and negative",
		input:    []int{3, -1, 4, -5, 2, 0},
		expected: []int{-5, -1, 0, 2, 3, 4},
	},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
			// Ensure original array is not modified
			if len(tt.input) > 0 {
				inputCopy := make([]int, len(tt.input))
				copy(inputCopy, tt.input)
				BubbleSort(tt.input)
				if !reflect.DeepEqual(tt.input, inputCopy) {
					t.Error("BubbleSort modified the original array")
				}
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectionSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SelectionSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertionSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range sortTestCases {
		t.Run(tt.name, func(t *testing.T) {
			result := HeapSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("HeapSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
