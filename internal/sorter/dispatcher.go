package sorter

import (
	"sort"

	"push-swap/internal/ops"
	"push-swap/internal/stack"
)

type Stack = stack.Stack

// Sort generates push-swap instructions to sort stack a (b starts empty).
// emit is called with each instruction name ("ra", "pb", etc.).
func Sort(a, b *Stack, emit func(string)) {
	n := a.Len()
	switch {
	case n <= 1:
		return
	case n <= 5:
		normalize(a)
		sortSmall(a, b, emit)
	default:
		normalize(a)
		if isSortedNormalized(a) {
			return
		}
		// Check for one-swap: sa
		ops.SwapA(a)
		if isSortedNormalized(a) {
			emit("sa")
			return
		}
		ops.SwapA(a) // undo swap

		// Check for one-rotation (ra or rra)
		pos := findMinPos(a)
		if pos != 0 {
			distTop := pos
			distBot := n - pos

			// Temporarily rotate to check if it's the only thing needed
			if distTop <= distBot {
				if isSortedAfterRotation(a, distTop, true) {
					for i := 0; i < distTop; i++ {
						applyAndEmit(a, b, "ra", emit)
					}
					return
				}
			} else {
				if isSortedAfterRotation(a, distBot, false) {
					for i := 0; i < distBot; i++ {
						applyAndEmit(a, b, "rra", emit)
					}
					return
				}
			}
		}

		sortChunk(a, b, emit, n)
	}
}

// normalize replaces each value in stack a with its rank (0 = smallest).
func normalize(a *Stack) {
	n := a.Len()
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		v, _ := a.Pop()
		vals[i] = v
	}
	sorted := make([]int, n)
	copy(sorted, vals)
	sort.Ints(sorted)
	rankOf := make(map[int]int, n)
	for r, v := range sorted {
		rankOf[v] = r
	}
	// Push in reverse so vals[0] (original top) ends up on top again.
	for i := n - 1; i >= 0; i-- {
		a.Push(rankOf[vals[i]])
	}
}

// findRankPos returns the logical index (0=top) of the element with the specific rank.
func findRankPos(s *Stack, rank int) int {
	for i := 0; i < s.Len(); i++ {
		if s.At(i) == rank {
			return i
		}
	}
	return -1
}

// isSortedNormalized checks if the stack is already in ideal order (0, 1, 2...).
func isSortedNormalized(a *Stack) bool {
	for i := 0; i < a.Len(); i++ {
		if a.At(i) != i {
			return false
		}
	}
	return true
}

// findMinPos returns the index of the element with rank 0.
func findMinPos(a *Stack) int {
	for i := 0; i < a.Len(); i++ {
		if a.At(i) == 0 {
			return i
		}
	}
	return 0
}

// isSortedAfterRotation checks if the stack would be sorted after a certain rotation.
func isSortedAfterRotation(a *Stack, dist int, forward bool) bool {
	n := a.Len()
	for i := 0; i < n; i++ {
		var actualIdx int
		if forward {
			actualIdx = (i + dist) % n
		} else {
			actualIdx = (i - dist + n) % n
		}
		if a.At(actualIdx) != i {
			return false
		}
	}
	return true
}

// applyAndEmit executes an instruction on the stacks and calls emit.
func applyAndEmit(a, b *Stack, instr string, emit func(string)) {
	switch instr {
	case "sa":
		ops.SwapA(a)
	case "sb":
		ops.SwapB(b)
	case "ss":
		ops.SwapBoth(a, b)
	case "pa":
		ops.PushA(a, b)
	case "pb":
		ops.PushB(a, b)
	case "ra":
		ops.RotateA(a)
	case "rb":
		ops.RotateB(b)
	case "rr":
		ops.RotateBoth(a, b)
	case "rra":
		ops.ReverseRotateA(a)
	case "rrb":
		ops.ReverseRotateB(b)
	case "rrr":
		ops.ReverseRotateBoth(a, b)
	}
	emit(instr)
}
