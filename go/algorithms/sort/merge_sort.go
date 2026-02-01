package sort

// MergeSort sorts an array using the merge sort algorithm
// Time Complexity: O(n log n)
// Space Complexity: O(n)
// Stable: Yes
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	mergeSortHelper(result, 0, len(result)-1)
	return result
}

func mergeSortHelper(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortHelper(arr, left, mid)
		mergeSortHelper(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}
