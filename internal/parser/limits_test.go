package parser

import "testing"

// Table-Driven Tests for validating MaxInt/MinInt boundaries.
func TestInt32Boundaries(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"max int32 accepted", []string{"2147483647"}, false},
		{"min int32 accepted", []string{"-2147483648"}, false},
		{"max int32 + 1 rejected (E03)", []string{"2147483648"}, true},
		{"min int32 - 1 rejected (E03)", []string{"-2147483649"}, true},
		{"all three boundary values", []string{"2147483647 -2147483648 0"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, exited := runParse(tt.args)
			if tt.wantErr && !exited {
				t.Errorf("expected error exit, got none")
			}
			if !tt.wantErr && exited {
				t.Errorf("unexpected error exit")
			}
		})
	}
}

// TestEC01EmptyWhitespace covers edge case EC01: empty or whitespace-only args.
func TestEC01EmptyWhitespace(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"empty string arg", []string{""}},
		{"whitespace only", []string{"   "}},
		{"tab only", []string{"\t"}},
		{"mixed whitespace", []string{"  \t  "}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, exited := runParse(tt.args)
			if !exited {
				t.Errorf("expected error exit for input %q, got none", tt.args)
			}
		})
	}
}
