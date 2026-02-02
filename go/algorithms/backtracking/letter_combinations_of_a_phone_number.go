package backtracking

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapping := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var result []string
	var backtrack func(index int, path []byte)
	backtrack = func(index int, path []byte) {
		if index == len(digits) {
			result = append(result, string(path))
			return
		}

		letters := mapping[digits[index]]
		for _, letter := range letters {
			backtrack(index+1, append(path, letter))
		}
	}

	backtrack(0, []byte{})
	return result

}
