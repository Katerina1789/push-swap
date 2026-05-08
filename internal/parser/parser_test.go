package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// runParse safely captures os.Exit calls for testing.
func runParse(args []string) (nums []int, exited bool) {
	// Replace os.Exit temporarily.
	old := ExitFunc
	defer func() { ExitFunc = old }()

	exited = false
	ExitFunc = func() {
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
		name string
		// golden is an optional file under testdata/golden/ whose content
		// must equal the compact result string ("error" or space-joined ints).
		golden  string
		args    []string
		want    []int
		wantErr bool
	}{
		{"valid single arg", "parse_valid.golden", []string{"1 2 3"}, []int{1, 2, 3}, false},
		{"valid multi args", "", []string{"1", "2", "3"}, []int{1, 2, 3}, false},
		{"empty string", "parse_ec01_error.golden", []string{""}, nil, true},
		{"whitespace only", "", []string{"   "}, nil, true},
		{"duplicate", "parse_e02_error.golden", []string{"1 2 2 3"}, nil, true},
		{"non-integer", "parse_e01_error.golden", []string{"0 one 2 3"}, nil, true},
		{"overflow", "parse_e03_error.golden", []string{"2147483648"}, nil, true},
		{"underflow", "", []string{"-2147483649"}, nil, true},
		{"no args", "", []string{}, []int{}, false},
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

			if tt.golden != "" {
				wantBytes, err := os.ReadFile(filepath.Join("testdata", "golden", tt.golden))
				if err != nil {
					t.Fatalf("missing golden file %s: %v", tt.golden, err)
				}
				wantStr := strings.TrimRight(string(wantBytes), "\r\n")

				var gotStr string
				if exited {
					gotStr = "error"
				} else {
					parts := make([]string, len(got))
					for i, n := range got {
						parts[i] = fmt.Sprintf("%d", n)
					}
					gotStr = strings.Join(parts, " ")
				}

				if gotStr != wantStr {
					t.Errorf("golden %s: got %q, want %q", tt.golden, gotStr, wantStr)
				}
			}
		})
	}
}

// FuzzParse ensures the parser never panics on arbitrary string input.
// Run with: go test -fuzz=FuzzParse ./internal/parser/
func FuzzParse(f *testing.F) {
	seeds := []string{
		"1 2 3",
		"",
		"   ",
		"2147483647 -2147483648 0",
		"2147483648",
		"-2147483649",
		"1 2 2 3",
		"abc",
		"0",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, input string) {
		args := []string{}
		if strings.TrimSpace(input) != "" {
			args = []string{input}
		}
		_, _ = runParse(args)
	})
}
