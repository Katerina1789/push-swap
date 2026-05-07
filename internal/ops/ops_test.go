package ops

import (
	"push-swap/internal/stack"
	"testing"
)

func stackOf(items ...int) *stack.Stack {
	s := stack.NewStack(len(items))
	for _, v := range items {
		s.Push(v)
	}
	return s
}

func readItems(s *stack.Stack) []int {
	temp := make([]int, 0, s.Len())
	for !s.IsEmpty() {
		v, _ := s.Pop()
		temp = append(temp, v)
	}

	items := make([]int, len(temp))
	for i := range temp {
		items[len(temp)-1-i] = temp[i]
	}

	for i := len(temp) - 1; i >= 0; i-- {
		s.Push(temp[i])
	}

	return items
}

func assertItems(t *testing.T, got *stack.Stack, want []int) {
	t.Helper()
	if got == nil {
		t.Fatalf("expected non-nil stack, got nil")
	}

	items := readItems(got)
	if len(items) != len(want) {
		t.Fatalf("stack length = %d, want %d", len(items), len(want))
	}
	for i := range items {
		if items[i] != want[i] {
			t.Fatalf("stack[%d] = %d, want %d", i, items[i], want[i])
		}
	}
}

func TestSwapA(t *testing.T) {
	tests := []struct {
		name  string
		input *stack.Stack
		want  []int
	}{
		{
			name:  "nil stack is no-op",
			input: nil,
			want:  nil,
		},
		{
			name:  "single element is no-op",
			input: stackOf(1),
			want:  []int{1},
		},
		{
			name:  "swap first two elements",
			input: stackOf(1, 2, 3),
			want:  []int{1, 3, 2},
		},
		{
			name:  "swap only top two, leave rest",
			input: stackOf(4, 5, 6, 7),
			want:  []int{4, 5, 7, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == nil {
				SwapA(nil)
				return
			}
			SwapA(tc.input)
			assertItems(t, tc.input, tc.want)
		})
	}
}

func TestSwapB(t *testing.T) {
	s := stackOf(8, 9)
	SwapB(s)
	assertItems(t, s, []int{9, 8})
}

func TestSwapBoth(t *testing.T) {
	a := stackOf(1, 2)
	b := stackOf(3, 4)

	SwapBoth(a, b)

	assertItems(t, a, []int{2, 1})
	assertItems(t, b, []int{4, 3})
}

func TestPushA(t *testing.T) {
	tests := []struct {
		name  string
		a     *stack.Stack
		b     *stack.Stack
		wantA []int
		wantB []int
	}{
		{
			name:  "push from b to a when b has items",
			a:     stackOf(1, 2),
			b:     stackOf(3, 4),
			wantA: []int{1, 2, 4},
			wantB: []int{3},
		},
		{
			name:  "no-op when b is empty",
			a:     stackOf(1),
			b:     stackOf(),
			wantA: []int{1},
			wantB: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			PushA(tc.a, tc.b)
			assertItems(t, tc.a, tc.wantA)
			assertItems(t, tc.b, tc.wantB)
		})
	}
}

func TestPushB(t *testing.T) {
	a := stackOf(5, 6)
	b := stackOf(7)

	PushB(a, b)

	assertItems(t, a, []int{5})
	assertItems(t, b, []int{7, 6})
}

func TestRotateA(t *testing.T) {
	tests := []struct {
		name  string
		input *stack.Stack
		want  []int
	}{
		{
			name:  "nil stack is no-op",
			input: nil,
			want:  nil,
		},
		{
			name:  "single element is no-op",
			input: stackOf(1),
			want:  []int{1},
		},
		{
			name:  "rotate 3 elements",
			input: stackOf(1, 2, 3),
			want:  []int{3, 1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == nil {
				RotateA(nil)
				return
			}
			RotateA(tc.input)
			assertItems(t, tc.input, tc.want)
		})
	}
}

func TestRotateB(t *testing.T) {
	s := stackOf(1, 2, 3)
	RotateB(s)
	assertItems(t, s, []int{3, 1, 2})
}

func TestRotateBoth(t *testing.T) {
	a := stackOf(1, 2)
	b := stackOf(3, 4)

	RotateBoth(a, b)

	assertItems(t, a, []int{2, 1})
	assertItems(t, b, []int{4, 3})
}

func TestReverseRotateA(t *testing.T) {
	tests := []struct {
		name  string
		input *stack.Stack
		want  []int
	}{
		{
			name:  "nil stack is no-op",
			input: nil,
			want:  nil,
		},
		{
			name:  "single element is no-op",
			input: stackOf(9),
			want:  []int{9},
		},
		{
			name:  "reverse rotate 4 elements",
			input: stackOf(1, 2, 3, 4),
			want:  []int{2, 3, 4, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.input == nil {
				ReverseRotateA(nil)
				return
			}
			ReverseRotateA(tc.input)
			assertItems(t, tc.input, tc.want)
		})
	}
}

func TestReverseRotateB(t *testing.T) {
	s := stackOf(1, 2, 3)
	ReverseRotateB(s)
	assertItems(t, s, []int{2, 3, 1})
}

func TestReverseRotateBoth(t *testing.T) {
	a := stackOf(1, 2, 3)
	b := stackOf(4, 5, 6)

	ReverseRotateBoth(a, b)

	assertItems(t, a, []int{2, 3, 1})
	assertItems(t, b, []int{5, 6, 4})
}
