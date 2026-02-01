package sort

// SelectionSort sorts an array using the selection sort algorithm
// Time Complexity: O(n²)
// Space Complexity: O(1)
// Stable: No
func SelectionSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			result[i], result[minIdx] = result[minIdx], result[i]
		}
	}

	return result
}
