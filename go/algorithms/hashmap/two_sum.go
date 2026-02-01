package hashmap

func twoSum(nums []int, target int) []int {
	hasher := make(map[int]int)

	for i, value := range nums {
		if v, ok := hasher[value]; ok {
			return []int{v, i}
		}
		hasher[target-value] = i
	}

	return []int{}
}
