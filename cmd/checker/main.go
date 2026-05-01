// Checker for validating the sequence of operations in the push_swap algorithm.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"push-swap/internal/errs"
	"push-swap/internal/ops"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func main() {
	nums := parser.Parse(os.Args[1:])

	// C07: no args → no output.
	if len(nums) == 0 {
		return
	}

	stackA := stack.NewStack(len(nums))
	stackB := stack.NewStack(len(nums))

	// Push in reverse so nums[0] ends up on top.
	for i := len(nums) - 1; i >= 0; i-- {
		stackA.Push(nums[i])
	}

	dispatchMap := map[string]func(*stack.Stack, *stack.Stack){
		"sa":  func(a, b *stack.Stack) { ops.SwapA(a) },
		"sb":  func(a, b *stack.Stack) { ops.SwapB(b) },
		"ss":  func(a, b *stack.Stack) { ops.SwapBoth(a, b) },
		"pa":  func(a, b *stack.Stack) { ops.PushA(a, b) },
		"pb":  func(a, b *stack.Stack) { ops.PushB(a, b) },
		"ra":  func(a, b *stack.Stack) { ops.RotateA(a) },
		"rb":  func(a, b *stack.Stack) { ops.RotateB(b) },
		"rr":  func(a, b *stack.Stack) { ops.RotateBoth(a, b) },
		"rra": func(a, b *stack.Stack) { ops.ReverseRotateA(a) },
		"rrb": func(a, b *stack.Stack) { ops.ReverseRotateB(b) },
		"rrr": func(a, b *stack.Stack) { ops.ReverseRotateBoth(a, b) },
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// Reject instructions with trailing/leading spaces.
		if line != strings.TrimSpace(line) {
			errs.ExitWithError()
		}

		fn, ok := dispatchMap[line]
		if !ok {
			errs.ExitWithError()
		}
		fn(stackA, stackB)
	}

	if err := scanner.Err(); err != nil {
		errs.ExitWithError()
	}

	if isSortedStack(stackA) && stackB.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// isSortedStack checks that stack a is sorted ascending from top to bottom
// (i.e. At(0) < At(1) < ... < At(n-1)).
func isSortedStack(s *stack.Stack) bool {
	n := s.Len()
	for i := 0; i < n-1; i++ {
		if s.At(i) >= s.At(i+1) {
			return false
		}
	}
	return true
}
