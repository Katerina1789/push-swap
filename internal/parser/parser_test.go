package parser

import (
	"push-swap/internal/errs"
	"testing"
)

// runParse safely captures os.Exit calls for testing.
func runParse(args []string) (nums []int, exited bool) {
	// Replace os.Exit temporarily.
	old := errs.ExitFunc
	defer func() { errs.ExitFunc = old }()

	exited = false
	errs.ExitFunc = func() {
		exited = true
	}

	nums = Parse(args)
	return nums, exited
}

// equalSlices compares two int slices.
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []int
		wantErr bool
	}{
		{"valid single arg", []string{"1 2 3"}, []int{1, 2, 3}, false},
		{"valid multi args", []string{"1", "2", "3"}, []int{1, 2, 3}, false},
		{"empty string", []string{""}, nil, true},
		{"whitespace only", []string{"   "}, nil, true},
		{"duplicate", []string{"1 2 2 3"}, nil, true},
		{"non-integer", []string{"1 two 3"}, nil, true},
		{"overflow", []string{"2147483648"}, nil, true},
		{"underflow", []string{"-2147483649"}, nil, true},
		{"no args", []string{}, []int{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, exited := runParse(tt.args)

			if tt.wantErr && !exited {
				t.Errorf("expected error exit, got none")
			}
			if !tt.wantErr && exited {
				t.Errorf("unexpected error exit")
			}

			if !tt.wantErr && !equalSlices(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
