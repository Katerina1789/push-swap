package sorter

import (
	"push-swap/internal/ops"
)

// sortChunk implements the "Butterfly" algorithm.
// It partitions A into B using a sliding window to minimize rotations,
// then pushes back to A in perfect order.
func sortChunk(a, b *Stack, emit func(string), n int) {
	chunkSize := calculateChunkSize(n)

	// Part 1: Partition A into B.
	// i tracks the number of elements already in B.
	for i := 0; a.Len() > 0; {
		topVal := a.At(0)
		if topVal <= i {
			// Element belongs to the "lower" part of the current window.
			// We push it to B and rotate it to the bottom of B.
			ops.PushB(a, b)
			emit("pb")
			ops.RotateB(b)
			emit("rb")
			i++
		} else if topVal <= i+chunkSize {
			// Element belongs to the "upper" part of the current window.
			// We push it to B and keep it at the top.
			ops.PushB(a, b)
			emit("pb")
			i++
		} else {
			// Element is not in the window yet, rotate A to find one.
			ops.RotateA(a)
			emit("ra")
		}
	}

	// Part 2: Push back from B to A.
	// We always look for the current maximum rank to ensure A ends up sorted.
	for b.Len() > 0 {
		target := b.Len() - 1
		pos := findRankPos(b, target)

		// Bring the target element to the top of B.
		moveToTopB(b, pos, emit)

		// Push to A.
		ops.PushA(a, b)
		emit("pa")
	}
}

// moveToTopB brings an element at pos to the top of stack b using the shortest path.
func moveToTopB(b *Stack, pos int, emit func(string)) {
	size := b.Len()
	if pos <= size/2 {
		for i := 0; i < pos; i++ {
			ops.RotateB(b)
			emit("rb")
		}
	} else {
		for i := 0; i < size-pos; i++ {
			ops.ReverseRotateB(b)
			emit("rrb")
		}
	}
}

func calculateChunkSize(n int) int {
	if n <= 100 {
		return 16
	}
	return 32
}
