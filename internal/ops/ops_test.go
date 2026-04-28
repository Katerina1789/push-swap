package operations

import (
    "testing"
)

// stack creates a new Stack instance with the provided items.
// It also makes a copy of the slice to avoid test mutation side effects.
func stack(items ...int) *Stack {
    return &Stack{Items: append([]int{}, items...)}
}

// assertItems compares the expected contents with the actual stack contents.
// It treats nil and empty slices as equivalent because a stack with zero items
// should behave the same whether the underlying slice is nil or empty.
func assertItems(t *testing.T, got *Stack, want []int) {
    t.Helper()
    if got == nil {
        t.Fatalf("expected non-nil stack, got nil")
    }
    // Compare lengths first
    if len(got.Items) != len(want) {
        t.Fatalf("stack length = %d, want %d", len(got.Items), len(want))
    }
    // Compare each element
    for i := range got.Items {
        if got.Items[i] != want[i] {
            t.Fatalf("stack[%d] = %d, want %d", i, got.Items[i], want[i])
        }
    }
}

func TestSwapA(t *testing.T) {
    tests := []struct {
        name string
        input *Stack
        want  []int
    }{
        // Edge case: SwapA should gracefully handle nil input without panicking.
        {
            name: "nil stack is no-op",
            input: nil,
            want: nil,
        },
        // Edge case: SwapA should do nothing if stack has fewer than 2 elements.
        {
            name: "single element is no-op",
            input: stack(1),
            want: []int{1},
        },
        // Core behavior: SwapA swaps the top two elements while preserving the rest.
        {
            name: "swap first two elements",
            input: stack(1, 2, 3),
            want: []int{2, 1, 3},
        },
        // Verification: Only the top two are swapped; deeper elements remain unchanged.
        {
            name: "swap only top two, leave rest",
            input: stack(4, 5, 6, 7),
            want: []int{5, 4, 6, 7},
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
    // SwapB is implemented by reusing the logic of SwapA.
    // We only need to verify that it behaves identically for stack b.
    s := stack(8, 9)
    SwapB(s)
    assertItems(t, s, []int{9, 8})
}

func TestSwapBoth(t *testing.T) {
    // SwapBoth performs SwapA and SwapB together.
    a := stack(1, 2)
    b := stack(3, 4)

    SwapBoth(a, b)

    assertItems(t, a, []int{2, 1})
    assertItems(t, b, []int{4, 3})
}

func TestPushA(t *testing.T) {
    tests := []struct {
        name string
        a    *Stack
        b    *Stack
        wantA []int
        wantB []int
    }{
        // Core behavior: PushA moves top element from b to top of a.
        {
            name: "push from b to a when b has items",
            a: stack(1, 2),
            b: stack(3, 4),
            wantA: []int{3, 1, 2},
            wantB: []int{4},
        },
        // Edge case: PushA does nothing if b is empty; a remains unchanged.
        {
            name: "no-op when b is empty",
            a: stack(1),
            b: stack(),
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
    // PushB moves the top element of stack a onto stack b.
    a := stack(5, 6)
    b := stack(7)

    PushB(a, b)

    assertItems(t, a, []int{6})
    assertItems(t, b, []int{5, 7})
}

func TestRotateA(t *testing.T) {
    tests := []struct {
        name string
        input *Stack
        want  []int
    }{
        // Edge case: RotateA should gracefully handle nil input without panicking.
        {
            name: "nil stack is no-op",
            input: nil,
            want: nil,
        },
        // Edge case: RotateA should do nothing if stack has fewer than 2 elements.
        {
            name: "single element is no-op",
            input: stack(1),
            want: []int{1},
        },
        // Core behavior: RotateA shifts all elements up; first becomes last.
        {
            name: "rotate 3 elements",
            input: stack(1, 2, 3),
            want: []int{2, 3, 1},
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
    // RotateB is identical to RotateA but applied to stack b.
    s := stack(1, 2, 3)
    RotateB(s)
    assertItems(t, s, []int{2, 3, 1})
}

func TestRotateBoth(t *testing.T) {
    // RotateBoth performs RotateA and RotateB simultaneously on both stacks.
    a := stack(1, 2)
    b := stack(3, 4)

    RotateBoth(a, b)

    assertItems(t, a, []int{2, 1})
    assertItems(t, b, []int{4, 3})
}

func TestReverseRotateA(t *testing.T) {
    tests := []struct {
        name string
        input *Stack
        want  []int
    }{
        // Edge case: ReverseRotateA should gracefully handle nil input without panicking.
        {
            name: "nil stack is no-op",
            input: nil,
            want: nil,
        },
        // Edge case: ReverseRotateA should do nothing if stack has fewer than 2 elements.
        {
            name: "single element is no-op",
            input: stack(9),
            want: []int{9},
        },
        // Core behavior: ReverseRotateA shifts down all elements; last becomes first.
        {
            name: "reverse rotate 4 elements",
            input: stack(1, 2, 3, 4),
            want: []int{4, 1, 2, 3},
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
    // ReverseRotateB is identical to ReverseRotateA but applied to stack b.
    s := stack(1, 2, 3)
    ReverseRotateB(s)
    assertItems(t, s, []int{3, 1, 2})
}

func TestReverseRotateBoth(t *testing.T) {
    // ReverseRotateBoth performs ReverseRotateA and ReverseRotateB simultaneously on both stacks.
    a := stack(1, 2, 3)
    b := stack(4, 5, 6)

    ReverseRotateBoth(a, b)

    assertItems(t, a, []int{3, 1, 2})
    assertItems(t, b, []int{6, 4, 5})
}
