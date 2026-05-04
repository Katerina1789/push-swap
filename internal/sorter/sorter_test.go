package sorter_test

import (
	"math/rand"
	"testing"

	"push-swap/internal/ops"
	"push-swap/internal/sorter"
	"push-swap/internal/stack"
)

// ── helpers ───────────────────────────────────────────────────────────────────

// collectOps initialises stacks, runs Sort, and returns the emitted instruction list.
func collectOps(input []int) []string {
	n := len(input)
	a := stack.NewStack(n)
	b := stack.NewStack(n)
	for i := n - 1; i >= 0; i-- {
		a.Push(input[i])
	}
	var out []string
	sorter.Sort(a, b, func(op string) { out = append(out, op) })
	return out
}

// applyOp dispatches a single push-swap instruction to the ops package.
func applyOp(op string, a, b *stack.Stack) {
	switch op {
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
}

// assertValid replays instructions on a fresh copy of input, then verifies:
//   - stack A is sorted in ascending order (top = smallest)
//   - stack B is empty
func assertValid(t *testing.T, input []int, instructions []string) {
	t.Helper()
	n := len(input)
	a := stack.NewStack(n)
	b := stack.NewStack(n)
	for i := n - 1; i >= 0; i-- {
		a.Push(input[i])
	}
	for _, op := range instructions {
		applyOp(op, a, b)
	}
	if !b.IsEmpty() {
		t.Errorf("stack B not empty after sort (input=%v)", input)
		return
	}
	for i := 0; i < a.Len()-1; i++ {
		if a.At(i) > a.At(i+1) {
			t.Errorf("stack A unsorted at [%d]=%d > [%d]=%d (input=%v)",
				i, a.At(i), i+1, a.At(i+1), input)
			return
		}
	}
}

// permutations returns all n! permutations of [0..n-1] via Heap's algorithm.
func permutations(n int) [][]int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	var result [][]int
	var gen func(k int)
	gen = func(k int) {
		if k == 1 {
			tmp := make([]int, n)
			copy(tmp, perm)
			result = append(result, tmp)
			return
		}
		for i := 0; i < k; i++ {
			gen(k - 1)
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}
		}
	}
	gen(n)
	return result
}

// randPerm returns a deterministic shuffle of [1..n] using the given seed.
func randPerm(n int, seed int64) []int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i + 1
	}
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(n, func(i, j int) { perm[i], perm[j] = perm[j], perm[i] })
	return perm
}

// ── edge / boundary cases ─────────────────────────────────────────────────────

// EC02: single element → no instructions needed.
func TestSortSingleElement(t *testing.T) {
	result := collectOps([]int{42})
	if len(result) != 0 {
		t.Errorf("single element: got %d ops, want 0 — ops=%v", len(result), result)
	}
}

// EC03 / C03: already-sorted input → no instructions needed.
func TestSortAlreadySorted(t *testing.T) {
	cases := [][]int{
		{1, 2},
		{1, 2, 3},
		{0, 1, 2, 3, 4, 5},
		{10, 20, 30, 40, 50},
		{-5, -3, 0, 7, 100},
	}
	for _, input := range cases {
		result := collectOps(input)
		if len(result) != 0 {
			t.Errorf("already sorted %v: got %d ops, want 0", input, len(result))
		}
	}
}

// EC04: values at int32 boundaries must be sorted correctly.
func TestSortIntBounds(t *testing.T) {
	input := []int{2147483647, -2147483648, 0}
	result := collectOps(input)
	assertValid(t, input, result)
}

// ── small-sort unit tests (n = 2 .. 6) ───────────────────────────────────────

func TestSortN2(t *testing.T) {
	for _, perm := range permutations(2) {
		result := collectOps(perm)
		assertValid(t, perm, result)
	}
}

// TASK-05: n=3 must sort in ≤ 2 instructions.
func TestSortN3(t *testing.T) {
	const limit = 2
	for _, perm := range permutations(3) {
		result := collectOps(perm)
		assertValid(t, perm, result)
		if len(result) > limit {
			t.Errorf("n=3 perm=%v: %d ops > limit %d", perm, len(result), limit)
		}
	}
}

func TestSortN4(t *testing.T) {
	for _, perm := range permutations(4) {
		result := collectOps(perm)
		assertValid(t, perm, result)
	}
}

// TASK-05 / C06 / P01: all 120 permutations of n=5 must sort in < 12 instructions.
func TestSortN5_AuditC06(t *testing.T) {
	const limit = 12
	for _, perm := range permutations(5) {
		result := collectOps(perm)
		assertValid(t, perm, result)
		if len(result) >= limit {
			t.Errorf("n=5 perm=%v: %d ops >= limit %d", perm, len(result), limit)
		}
	}
}

// n=6 is the upper boundary of the small-sort path (BFS lookup table).
func TestSortN6(t *testing.T) {
	for _, perm := range permutations(6) {
		result := collectOps(perm)
		assertValid(t, perm, result)
	}
}

// ── audit-case unit tests ─────────────────────────────────────────────────────

// C02: specific 6-element input must produce a valid sort in < 9 instructions.
func TestAuditC02(t *testing.T) {
	input := []int{2, 1, 3, 6, 5, 8}
	const limit = 9
	result := collectOps(input)
	assertValid(t, input, result)
	if len(result) >= limit {
		t.Errorf("C02 [2,1,3,6,5,8]: %d ops >= limit %d — ops=%v", len(result), limit, result)
	}
}

// ── fast-path / optimisation checks ──────────────────────────────────────────

// Two-element swap: the only valid instruction is exactly [sa].
func TestSort_OneSwap(t *testing.T) {
	input := []int{5, 3}
	result := collectOps(input)
	if len(result) != 1 || result[0] != "sa" {
		t.Errorf("two-element swap: got %v, want [sa]", result)
	}
}

// n=7..20 exercises the small→chunk boundary and the chunk algorithm on tiny inputs.
func TestSort_ChunkBoundary(t *testing.T) {
	for n := 7; n <= 20; n++ {
		perm := randPerm(n, int64(n))
		result := collectOps(perm)
		assertValid(t, perm, result)
	}
}

// ── large-sort instruction-count tests ───────────────────────────────────────

// P02 / TASK-06 / TASK-09: n=100 must sort in < 700 instructions across 5 seeds.
func TestSort100_Limit(t *testing.T) {
	const limit = 700
	seeds := []int64{1, 42, 100, 777, 12345}
	for _, seed := range seeds {
		perm := randPerm(100, seed)
		result := collectOps(perm)
		assertValid(t, perm, result)
		if len(result) >= limit {
			t.Errorf("n=100 seed=%d: %d ops >= limit %d", seed, len(result), limit)
		}
	}
}

// P03 / TASK-06 / TASK-09: n=500 must sort in < 5500 instructions across 5 seeds.
func TestSort500_Limit(t *testing.T) {
	const limit = 5500
	seeds := []int64{1, 42, 100, 777, 12345}
	for _, seed := range seeds {
		perm := randPerm(500, seed)
		result := collectOps(perm)
		assertValid(t, perm, result)
		if len(result) >= limit {
			t.Errorf("n=500 seed=%d: %d ops >= limit %d", seed, len(result), limit)
		}
	}
}

// ── benchmarks ────────────────────────────────────────────────────────────────

func BenchmarkSort100(b *testing.B) {
	perm := randPerm(100, 42)
	n := len(perm)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := stack.NewStack(n)
		bk := stack.NewStack(n)
		for j := n - 1; j >= 0; j-- {
			a.Push(perm[j])
		}
		b.StartTimer()
		sorter.Sort(a, bk, func(string) {})
	}
}

func BenchmarkSort500(b *testing.B) {
	perm := randPerm(500, 42)
	n := len(perm)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := stack.NewStack(n)
		bk := stack.NewStack(n)
		for j := n - 1; j >= 0; j-- {
			a.Push(perm[j])
		}
		b.StartTimer()
		sorter.Sort(a, bk, func(string) {})
	}
}
