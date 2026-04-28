package parser

import "strconv"

// hasDuplicates checks for duplicate integers.
func hasDuplicates(nums []int) bool {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}

// isValidInt32 checks if a string fits in int32 range.
func isValidInt32(s string) bool {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return false
	}
	return v >= -2147483648 && v <= 2147483647
}
