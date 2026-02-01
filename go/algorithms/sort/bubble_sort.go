package sort

// BubbleSort sorts an array using the bubble sort algorithm
// Time Complexity: O(n²) worst/average case, O(n) best case
// Space Complexity: O(1)
// Stable: Yes
func BubbleSort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		// If no swaps occurred, array is sorted
		if !swapped {
			break
		}
	}

	return result
}
