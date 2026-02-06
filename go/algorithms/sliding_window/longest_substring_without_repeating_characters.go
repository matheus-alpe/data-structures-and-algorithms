package slidingwindow

func lengthOfLongestSubstring(s string) int {
	left := 0
	result := 0
	frequency := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if value, ok := frequency[s[right]]; ok && value >= left {
			left = value + 1
		}

		result = max(result, right-left+1)
		frequency[s[right]] = right
	}

	return result
}
