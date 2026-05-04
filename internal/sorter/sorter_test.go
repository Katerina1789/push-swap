package sorter_test

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	"push-swap/internal/ops"
	"push-swap/internal/sorter"
	"push-swap/internal/stack"
)

// collectOps initializes the stacks with input data, executes the Sort algorithm,
// and returns the sequence of generated instructions.
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

// applyOp executes a specific push-swap instruction on the provided stacks.
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

// assertValid replays the generated instructions on a new stack containing the original input.
// It verifies that stack A is correctly sorted in ascending order and that
// stack B is empty upon completion.
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

// ExampleSort demonstrates how to use the Sort function.
func ExampleSort() {
	a := stack.NewStack(3)
	b := stack.NewStack(3)
	// Push values so that top is 2, then 1, then 3
	a.Push(3)
	a.Push(1)
	a.Push(2)
	sorter.Sort(a, b, func(op string) { fmt.Println(op) })
	// Output:
	// sa
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

// TestSort_TableDriven validates the sorter against specific edge cases and audit requirements.
func TestSort_TableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		limit int // Optional limit for instruction count
	}{
		{"EC02 Single element", []int{42}, 0},
		{"EC03 Already sorted 2", []int{1, 2}, 0},
		{"EC03 Already sorted 6", []int{0, 1, 2, 3, 4, 5}, 0},
		{"EC04 Int bounds", []int{2147483647, -2147483648, 0}, 10},
		{"C02 Audit case n=6", []int{2, 1, 3, 6, 5, 8}, 9},
		{"Fast path: One swap n=7", []int{1, 0, 2, 3, 4, 5, 6}, 2},
		{"Fast path: One rotation n=7", []int{1, 2, 3, 4, 5, 6, 0}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := collectOps(tt.input)
			assertValid(t, tt.input, result)
			if tt.limit >= 0 && len(result) > tt.limit {
				t.Errorf("got %d ops, want <= %d", len(result), tt.limit)
			}
		})
	}
}

// TestSort_SmallPermutations exhaustively tests all permutations for small stack sizes (n=2 to n=6).
// These stack sizes are handled by optimal BFS-generated lookup tables, as required by
// Audit cases C02 and C06.
func TestSort_SmallPermutations(t *testing.T) {
	limits := map[int]int{
		2: 1,
		3: 2,
		4: 12,
		5: 12,
		6: 15,
	}

	for n := 2; n <= 6; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			// We don't use t.Parallel() here because generating 720 perms for n=6
			// inside parallel tests can consume significant memory.
			perms := permutations(n)
			limit := limits[n]
			for _, p := range perms {
				result := collectOps(p)
				assertValid(t, p, result)
				if len(result) > limit {
					t.Fatalf("n=%d perm=%v: got %d ops, want <= %d", n, p, len(result), limit)
				}
			}
		})
	}
}

// TestSort_ChunkBoundary verifies behavior at the transition boundary to the chunk-based algorithm.
func TestSort_ChunkBoundary(t *testing.T) {
	for n := 7; n <= 20; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			t.Parallel()
			perm := randPerm(n, int64(n))
			result := collectOps(perm)
			assertValid(t, perm, result)
		})
	}
}

// TestSort_Limit100 verifies that the algorithm sorts 100 elements within the audit limit.
// This covers performance requirements P02 and C12, ensuring instruction counts
// remain below 700 across various random seeds.
func TestSort_Limit100(t *testing.T) {
	t.Parallel()
	const limit = 700
	seeds := []int64{1, 42, 100, 777, 12345}
	for _, seed := range seeds {
		t.Run(fmt.Sprintf("seed=%d", seed), func(t *testing.T) {
			t.Parallel()
			perm := randPerm(100, seed)
			result := collectOps(perm)
			assertValid(t, perm, result)
			if len(result) >= limit {
				t.Errorf("n=100 seed=%d: %d ops >= limit %d", seed, len(result), limit)
			}
		})
	}
}

// TestSort_Limit500 verifies that the algorithm sorts 500 elements within the audit limit.
// This test covers performance requirements P03 and C14, ensuring instruction counts
// remain below 5500 across various random seeds.
func TestSort_Limit500(t *testing.T) {
	t.Parallel()
	const limit = 5500
	seeds := []int64{1, 42, 100, 777, 12345}
	for _, seed := range seeds {
		t.Run(fmt.Sprintf("seed=%d", seed), func(t *testing.T) {
			t.Parallel()
			perm := randPerm(500, seed)
			result := collectOps(perm)
			assertValid(t, perm, result)
			if len(result) >= limit {
				t.Errorf("n=500 seed=%d: %d ops >= limit %d", seed, len(result), limit)
			}
		})
	}
}

// FuzzSort explores random permutations to ensure the sorter never panics
// and always produces a valid instruction set for any unique integer sequence.
func FuzzSort(f *testing.F) {
	// Seed with known interesting permutations
	f.Add([]byte{0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00})
	f.Add([]byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00})

	f.Fuzz(func(t *testing.T, data []byte) {
		// We need at least one int (4 bytes) and at most 100 for performance stability
		if len(data) < 4 || len(data) > 400 || len(data)%4 != 0 {
			t.Skip()
		}
		nums := make([]int, 0, len(data)/4)
		seen := make(map[int]bool)
		for i := 0; i < len(data); i += 4 {
			val := int(int32(binary.LittleEndian.Uint32(data[i : i+4])))
			if seen[val] {
				t.Skip() // Push-swap requires unique integers; skip duplicates
			}
			seen[val] = true
			nums = append(nums, val)
		}

		result := collectOps(nums)
		assertValid(t, nums, result)
	})
}

// BenchmarkSort100 measures the execution time and memory allocation for sorting 100 elements.
func BenchmarkSort100(b *testing.B) {
	perm := randPerm(100, 42)
	n := len(perm)
	b.ReportAllocs()
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

// BenchmarkSort500 measures the execution time and memory allocation for sorting 500 elements.
func BenchmarkSort500(b *testing.B) {
	perm := randPerm(500, 42)
	n := len(perm)
	b.ReportAllocs()
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
