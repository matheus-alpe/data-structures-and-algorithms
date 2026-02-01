package sort

// QuickSort sorts an array using the quicksort algorithm
// Time Complexity: O(n log n) average, O(n²) worst case
// Space Complexity: O(log n)
// Stable: No
func QuickSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
