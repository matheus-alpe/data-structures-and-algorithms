package sort

import (
	"testing"
)

// Benchmark tests to compare algorithm performance
func BenchmarkBubbleSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		BubbleSort(input)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		SelectionSort(input)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		InsertionSort(input)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		QuickSort(input)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		MergeSort(input)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50}
	for i := 0; i < b.N; i++ {
		HeapSort(input)
	}
}
