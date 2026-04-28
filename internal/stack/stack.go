package stack

// Stack holds a slice of ints representing a LIFO stack.
type Stack struct {
	data []int
}

// NewStack creates a stack with pre-allocated capacity.
func NewStack(capacity int) *Stack {
	return &Stack{
		data: make([]int, 0, capacity),
	}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element.
func (s *Stack) Pop() (int, bool) {
	n := len(s.data)
	if n == 0 {
		return 0, false
	}
	v := s.data[n-1]
	s.data = s.data[:n-1]
	return v, true
}

// Peek returns the top element without removing it.
func (s *Stack) Peek() (int, bool) {
	n := len(s.data)
	if n == 0 {
		return 0, false
	}
	return s.data[n-1], true
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Len returns the number of elements in the stack.
func (s *Stack) Len() int {
	return len(s.data)
}
