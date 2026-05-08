// Push_swap algorithm implementation.
package main

import (
	"fmt"
	"os"

	"push-swap/internal/parser"
	"push-swap/internal/sorter"
	"push-swap/internal/stack"
)

func main() {
	args := os.Args
	if len(args) > 0 {
		args = args[1:]
	}
	nums := parser.Parse(args)

	if len(nums) <= 1 || isSorted(nums) {
		return
	}

	n := len(nums)
	a := stack.NewStack(n)
	b := stack.NewStack(n)

	// Push in reverse so nums[0] ends up on top.
	for i := n - 1; i >= 0; i-- {
		a.Push(nums[i])
	}

	sorter.Sort(a, b, func(s string) {
		fmt.Println(s)
	})
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
