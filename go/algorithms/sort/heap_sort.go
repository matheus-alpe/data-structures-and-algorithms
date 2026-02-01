package sort

// HeapSort sorts an array using the heap sort algorithm
// Time Complexity: O(n log n)
// Space Complexity: O(1)
// Stable: No
func HeapSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}

	return result
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
