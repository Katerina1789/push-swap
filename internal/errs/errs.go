// Centralized error handling for the application.
// This package defines custom error types and helper functions
// for creating and managing errors throughout the application.

package errs

import (
	"fmt"
	"os"
)

// ExitFunc allows tests to override os.Exit.
var ExitFunc = func() {
	os.Exit(1)
}

// ExitWithError prints the mandatory error message to stderr and exits with status 1.
func ExitWithError() {
	fmt.Fprintln(os.Stderr, "Error")
	ExitFunc()
}
