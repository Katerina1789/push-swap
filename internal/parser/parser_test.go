package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"push-swap/internal/errs"
	"strings"
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
		// golden is an optional file under testdata/golden/ whose content
		// must equal the compact result string ("error" or space-joined ints).
		golden string
	}{
		{"valid single arg", []string{"1 2 3"}, []int{1, 2, 3}, false, "parse_valid.golden"},
		{"valid multi args", []string{"1", "2", "3"}, []int{1, 2, 3}, false, ""},
		{"empty string", []string{""}, nil, true, "parse_ec01_error.golden"},
		{"whitespace only", []string{"   "}, nil, true, ""},
		{"duplicate", []string{"1 2 2 3"}, nil, true, "parse_e02_error.golden"},
		{"non-integer", []string{"0 one 2 3"}, nil, true, "parse_e01_error.golden"},
		{"overflow", []string{"2147483648"}, nil, true, "parse_e03_error.golden"},
		{"underflow", []string{"-2147483649"}, nil, true, ""},
		{"no args", []string{}, []int{}, false, ""},
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
