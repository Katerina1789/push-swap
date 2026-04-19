// Table-Driven Tests for the parser functionality.
// This is the most critical area for unit tests.
// Is needed to verify that it correctly catches:
// - Non-integer strings (e.g., "1a")
// - Duplicates (e.g., "1 2 2")
// - Integer overflows (e.g., values > MaxInt32)
// - Empty strings or just whitespace.