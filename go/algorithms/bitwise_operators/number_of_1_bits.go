package bitwiseoperators

func hammingWeight(num int) int {
	count := 0

	for num != 0 {
		count += num & 1
		num >>= 1
	}

	return count
}
