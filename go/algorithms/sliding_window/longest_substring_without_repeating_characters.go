package slidingwindow

func lengthOfLongestSubstring(s string) int {
	left := 0
	result := 0
	frequency := make([]int, 128)

	for right := 0; right < len(s); right++ {
		frequency[s[right]]++

		for frequency[s[right]] > 1 {
			frequency[s[left]]--
			left++
		}

		result = max(result, right-left+1)
	}

	return result
}
