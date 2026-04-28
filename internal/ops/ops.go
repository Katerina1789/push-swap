// Operator for defining the operations available in the push_swap algorithm
// (push, swap, rotate, reverse rotate).
package operations

// We assume the Stack structure looks like this. 
type Stack struct {
	Items []int
}

// SwapA swaps the first 2 elements at the top of stack a.
// If stack a has fewer than 2 elements, this operation does nothing.
func SwapA(a *Stack) {
	if a == nil || len(a.Items) < 2 {
		return
	}
	a.Items[0], a.Items[1] = a.Items[1], a.Items[0]
}

// SwapB swaps the first 2 elements at the top of stack b.
// If stack b has fewer than 2 elements, this operation does nothing.
func SwapB(b *Stack) {
	SwapA(b) // Reuse the same logic as SwapA
}

// SwapBoth performs SwapA and SwapB simultaneously on both stacks.
func SwapBoth(a *Stack, b *Stack) {
	SwapA(a)
	SwapB(b)
}

// PushA moves the top element from stack b to the top of stack a.
// If stack b is empty, this operation does nothing.
func PushA(a *Stack, b *Stack) {
	if b == nil || len(b.Items) == 0 {
		return
	}
	// Extract the top element from stack B
	topElement := b.Items[0]
	// Remove it from stack B
	b.Items = b.Items[1:]
	// Insert it at the top of stack A
	a.Items = append([]int{topElement}, a.Items...)
}

// PushB moves the top element from stack a to the top of stack b.
// If stack a is empty, this operation does nothing.
func PushB(a *Stack, b *Stack) {
	PushA(b, a) // Reuse the logic by swapping the arguments
}

// RotateA shifts all elements in stack a upward by 1 position.
// The element at the top (index 0) moves to the bottom.
func RotateA(a *Stack) {
	if a == nil || len(a.Items) < 2 {
		return
	}
	top := a.Items[0]
	a.Items = append(a.Items[1:], top)
}

// RotateB shifts all elements in stack b upward by 1 position.
// The element at the top (index 0) moves to the bottom.
func RotateB(b *Stack) {
	RotateA(b)
}

// RotateBoth performs RotateA and RotateB simultaneously on both stacks.
func RotateBoth(a *Stack, b *Stack) {
	RotateA(a)
	RotateB(b)
}

// ReverseRotateA shifts all elements in stack a downward by 1 position.
// The element at the bottom moves to the top.
func ReverseRotateA(a *Stack) {
	if a == nil || len(a.Items) < 2 {
		return
	}
	lastIdx := len(a.Items) - 1
	bottom := a.Items[lastIdx]
	a.Items = append([]int{bottom}, a.Items[:lastIdx]...)
}

// ReverseRotateB shifts all elements in stack b downward by 1 position.
// The element at the bottom moves to the top.
func ReverseRotateB(b *Stack) {
	ReverseRotateA(b)
}

// ReverseRotateBoth performs ReverseRotateA and ReverseRotateB simultaneously on both stacks.
func ReverseRotateBoth(a *Stack, b *Stack) {
	ReverseRotateA(a)
	ReverseRotateB(b)
}