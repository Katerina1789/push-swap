package parser

import (
	"strconv"
	"strings"

	"push-swap/internal/errs"
)

// Parse converts CLI args into validated integers.
func Parse(args []string) []int {
	// No arguments → return empty slice (valid case).
	if len(args) == 0 {
		return []int{}
	}

	// Collect tokens from single-string or multi-arg formats.
	var tokens []string
	for _, arg := range args {
		fields := strings.Fields(arg)
		if len(fields) == 0 {
			errs.ExitWithError() // empty or whitespace-only argument
		}
		tokens = append(tokens, fields...)
	}

	// Convert tokens to integers with validation.
	nums := make([]int, 0, len(tokens))
	seen := make(map[int]bool)

	for _, tok := range tokens {
		// Validate int32 range.
		if !isValidInt32(tok) {
			errs.ExitWithError()
		}

		// Convert to int.
		n, err := strconv.Atoi(tok)
		if err != nil {
			errs.ExitWithError()
		}

		// Detect duplicates.
		if seen[n] {
			errs.ExitWithError()
		}
		seen[n] = true

		nums = append(nums, n)
	}

	return nums
}
