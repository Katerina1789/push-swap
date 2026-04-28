package stack

import (
	"testing"
)

// TestStackOperations covers Push, Pop, Peek using table-driven tests.
func TestStackOperations(t *testing.T) {
	tests := []struct {
		name    string
		run     func(*Stack) (int, bool)
		wantVal int
		wantOK  bool
	}{
		{
			name: "Push then Pop returns same value",
			run: func(s *Stack) (int, bool) {
				s.Push(5)
				return s.Pop()
			},
			wantVal: 5,
			wantOK:  true,
		},
		{
			name: "Pop on empty returns false",
			run: func(s *Stack) (int, bool) {
				return s.Pop()
			},
			wantVal: 0,
			wantOK:  false,
		},
		{
			name: "Peek on empty returns false",
			run: func(s *Stack) (int, bool) {
				return s.Peek()
			},
			wantVal: 0,
			wantOK:  false,
		},
		{
			name: "Push then Peek returns same value",
			run: func(s *Stack) (int, bool) {
				s.Push(42)
				return s.Peek()
			},
			wantVal: 42,
			wantOK:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{} // zero-value usable
			gotVal, gotOK := tt.run(s)

			if gotVal != tt.wantVal || gotOK != tt.wantOK {
				t.Errorf("got (%d, %v), want (%d, %v)",
					gotVal, gotOK, tt.wantVal, tt.wantOK)
			}
		})
	}
}

// TestZeroValueUsable ensures var s Stack works without initialization.
func TestZeroValueUsable(t *testing.T) {
	var s Stack
	if !s.IsEmpty() {
		t.Errorf("zero-value stack should be empty")
	}
	if s.Len() != 0 {
		t.Errorf("zero-value stack length should be 0")
	}
}

// TestCapacityHint verifies NewStack pre-allocates capacity.
func TestCapacityHint(t *testing.T) {
	s := NewStack(100)
	if cap(s.data) < 100 {
		t.Errorf("expected capacity >= 100, got %d", cap(s.data))
	}
}

// BenchmarkStack500 checks allocations for 500 push/pop operations.
func BenchmarkStack500(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewStack(500)
		for j := 0; j < 500; j++ {
			s.Push(j)
		}
		for j := 0; j < 500; j++ {
			s.Pop()
		}
	}
}
