// Checker for validating the sequence of operations in the push_swap algorithm.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"push-swap/internal/errs"
	"push-swap/internal/parser"
	"push-swap/internal/stack"
)

func main() {
	// Parse and validate input arguments.
	nums := parser.Parse(os.Args[1:])

	// Initialize stacks.
	stackA := stack.NewStack(len(nums))
	stackB := stack.NewStack(len(nums))

	// Push all numbers to stack A.
	for _, n := range nums {
		stackA.Push(n)
	}

	// Read instructions from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines at the end of input.
		if line == "" {
			continue
		}

		// Reject incorrectly formatted instructions (trailing spaces).
		if line != strings.TrimSpace(line) {
			errs.ExitWithError()
		}

		// TODO: Execute instruction (TASK-03)
		errs.ExitWithError()
	}

	if err := scanner.Err(); err != nil {
		errs.ExitWithError()
	}

	// Check if stack A is sorted and stack B is empty.
	if isSortedStack(stackA) && stackB.IsEmpty() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// isSortedStack checks if stack A is in ascending order.
func isSortedStack(s *stack.Stack) bool {
	// TODO: Implement proper sorted check.
	return true
}
