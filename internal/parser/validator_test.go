package parser

import "testing"

// TestHasDuplicates checks duplicate detection logic.
func TestHasDuplicates(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want bool
	}{
		{"no duplicates", []int{1, 2, 3}, false},
		{"with duplicates", []int{1, 2, 2, 3}, true},
		{"single element", []int{5}, false},
		{"empty slice", []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasDuplicates(tt.in)
			if got != tt.want {
				t.Errorf("hasDuplicates(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

// TestIsValidInt32 checks int32 boundary validation.
func TestIsValidInt32(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"valid small", "42", true},
		{"valid negative", "-42", true},
		{"max int32", "2147483647", true},
		{"min int32", "-2147483648", true},
		{"overflow", "2147483648", false},
		{"underflow", "-2147483649", false},
		{"non-integer", "abc", false},
		{"empty", "", false},
		{"whitespace", "   ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidInt32(tt.in)
			if got != tt.want {
				t.Errorf("isValidInt32(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
