// Push_swap algorithm implementation.
package main

import (
	"os"
	"push-swap/internal/parser"
)

func main() {
	// Parse and validate input arguments.
	nums := parser.Parse(os.Args[1:])

	// If already sorted or empty, output nothing.
	if len(nums) <= 1 || isSorted(nums) {
		return
	}

	// TODO: Implement sorting algorithm.
}

// isSorted checks if the slice is in ascending order.
func isSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
