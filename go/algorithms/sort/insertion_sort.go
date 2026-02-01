package sort

// InsertionSort sorts an array using the insertion sort algorithm
// Time Complexity: O(n²) worst/average case, O(n) best case
// Space Complexity: O(1)
// Stable: Yes
func InsertionSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 1; i < n; i++ {
		key := result[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}
