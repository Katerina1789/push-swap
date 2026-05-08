// Operator for defining the operations available in the push_swap algorithm
// (push, swap, rotate, reverse rotate).
package ops

import "push-swap/internal/stack"

type Stack = stack.Stack

// SwapA swaps the top 2 elements at the top of stack a.
// If stack a has fewer than 2 elements, this operation does nothing.
func SwapA(a *Stack) {
	if a == nil || a.Len() < 2 {
		return
	}

	top, _ := a.Pop()
	second, _ := a.Pop()
	a.Push(top)
	a.Push(second)
}

// SwapB swaps the first 2 elements at the top of stack b.
// If stack b has fewer than 2 elements, this operation does nothing.
func SwapB(b *Stack) {
	SwapA(b)
}

// SwapBoth performs SwapA and SwapB simultaneously on both stacks.
func SwapBoth(a *Stack, b *Stack) {
	SwapA(a)
	SwapB(b)
}

// PushA moves the top element from stack b to the top of stack a.
// If stack b is empty, this operation does nothing.
func PushA(a *Stack, b *Stack) {
	if b == nil || b.IsEmpty() {
		return
	}

	val, _ := b.Pop()
	a.Push(val)
}

// PushB moves the top element from stack a to the top of stack b.
// If stack a is empty, this operation does nothing.
func PushB(a *Stack, b *Stack) {
	PushA(b, a) // Reuse the logic by swapping the arguments
}

// RotateA shifts the top element of stack a to the bottom.
func RotateA(a *Stack) {
	if a == nil || a.Len() < 2 {
		return
	}

	top, ok := a.Pop()
	if !ok {
		return
	}
	buffer := make([]int, 0, a.Len())
	for !a.IsEmpty() {
		if val, ok := a.Pop(); ok {
			buffer = append(buffer, val)
		}
	}

	a.Push(top)
	for i := len(buffer) - 1; i >= 0; i-- {
		a.Push(buffer[i])
	}
}

// RotateB shifts the top element of stack b to the bottom.
func RotateB(b *Stack) {
	RotateA(b)
}

// RotateBoth performs RotateA and RotateB simultaneously on both stacks.
func RotateBoth(a *Stack, b *Stack) {
	RotateA(a)
	RotateB(b)
}

// ReverseRotateA shifts the bottom element of stack a to the top.
func ReverseRotateA(a *Stack) {
	if a == nil || a.Len() < 2 {
		return
	}

	buffer := make([]int, 0, a.Len())
	for a.Len() > 1 {
		if val, ok := a.Pop(); ok {
			buffer = append(buffer, val)
		}
	}

	bottom, ok := a.Pop()
	if !ok {
		return
	}
	for i := len(buffer) - 1; i >= 0; i-- {
		a.Push(buffer[i])
	}
	a.Push(bottom)
}

// ReverseRotateB shifts all elements in stack b to the top.
func ReverseRotateB(b *Stack) {
	ReverseRotateA(b)
}

// ReverseRotateBoth performs ReverseRotateA and ReverseRotateB simultaneously on both stacks.
func ReverseRotateBoth(a *Stack, b *Stack) {
	ReverseRotateA(a)
	ReverseRotateB(b)
}
