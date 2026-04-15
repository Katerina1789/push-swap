# Error Cases: push-swap

This document outlines the expected error handling behavior and validation rules for the `push-swap` and `checker` programs. It ensures the applications handle invalid inputs, duplicate numbers, and unknown instructions gracefully by outputting the strict error format required by the audit.

---

### Summary Table of Error Cases

| Case ID | Program | Description | Input Argument(s) / Stdin | Expected Output |
|:---:|---|---|---|---|
| E01 | push-swap/checker | Non-integer argument | `./push-swap "0 one 2 3"` | `Error\n` on stderr |
| E02 | push-swap/checker | Duplicate integer | `./push-swap "1 2 2 3"` | `Error\n` on stderr |
| E03 | push-swap/checker | Integer overflow/underflow | `./push-swap "2147483648"` | `Error\n` on stderr |
| E04 | checker | Invalid instruction | `echo -e "sa\nbadcmd\n" \| ./checker "1 2 3"` | `Error\n` on stderr |
| E05 | checker | Incorrectly formatted instruction | `echo -e "sa \n" \| ./checker "1 2 3"` | `Error\n` on stderr |

---

## Error Cases Detailed

### Error Case E01: Non-Integer Argument
**Description:**  
Tests the input validation when the user provides a string that cannot be parsed as a valid integer.

**Input Argument:**
```bash
./push-swap "0 one 2 3"
./checker "0 one 2 3"
```
**Validation Rules:**
- **Output:** Must print exactly `Error\n`.
- **Stream:** The output must be written to standard error (`stderr`), not standard output (`stdout`).
- **Exit Code:** Should exit with a non-zero status code.

---

### Error Case E02: Duplicate Integer
**Description:**  
Tests the core validation constraint that all numbers in the stack must be unique.

**Input Argument:**
```bash
./push-swap "1 2 2 3"
./checker "1 2 2 3"
```
**Validation Rules:**
- **Output:** Must print exactly `Error\n` on `stderr`.

---

### Error Case E03: Integer Overflow / Underflow
**Description:**  
Tests the parser's limits. The arguments must fit within standard integer bounds. Numbers exceeding `MaxInt` or `MinInt` (often tested with 32-bit limits depending on system, but definitely out-of-bounds strings) should fail safely.

**Input Argument:**
```bash
./push-swap "1 2 9999999999999999999999 3"
```
**Validation Rules:**
- **Output:** Must print exactly `Error\n` on `stderr`.

---

### Error Case E04: Invalid Instruction (checker)
**Description:**  
Tests the `checker` program's ability to identify and reject invalid or incorrectly formatted instructions passed via standard input.

**Input Argument / Trigger:**
```bash
echo -e "sa\nunknown_cmd\npb\n" | ./checker "3 2 1"
```
**Validation Rules:**
- **Output:** Must print exactly `Error\n` on `stderr`.
- **Execution:** The checker must immediately halt execution upon encountering the invalid instruction.

---

### Error Case E05: Incorrectly Formatted Instruction (checker)
**Description:**  
Tests the strictness of the `checker`'s instruction parser. Valid commands that contain trailing spaces, carriage returns (`\r`), or empty lines must be rejected as incorrectly formatted.

**Input Argument / Trigger:**
```bash
echo -e "sa \n\n" | ./checker "3 2 1"
```
**Validation Rules:**
- **Output:** Must print exactly `Error\n` on `stderr` and immediately halt execution.