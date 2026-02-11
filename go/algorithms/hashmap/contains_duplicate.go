package hashmap

func containsDuplicate(nums []int) bool {
	hasher := make(map[int]bool)
	for _, num := range nums {
		if hasher[num] {
			return true
		}
		hasher[num] = true
	}
	return false
}
