# Edge Cases: push-swap

This document outlines edge cases specific to the `push-swap` and `checker` programs. These scenarios test the robustness, adaptability, and boundaries of the core sorting logic and input parsing under unusual or extreme conditions. Error cases (like invalid formatting or duplicates) are covered in `error-cases.md`.

---

### Summary Table of Edge Cases

| Case ID | Category | Description | Input Argument(s) | Expected Behavior / Output |
|:---:|---|---|---|---|
| EC01 | Input Boundary | Empty or Whitespace String | `./push-swap ""` / `"   "` | `Error\n` on stderr. (Fails integer parsing). |
| EC02 | Input Boundary | Single Number | `./push-swap "42"` | Displays nothing (already sorted). |
| EC03 | Input Boundary | Already Sorted | `./push-swap "1 2 3 4"` | Displays nothing (0 instructions). |
| EC04 | Data Limits | Max/Min Int bounds | `./push-swap "2147483647 -2147483648 0"` | Parsed correctly and outputs valid sorting instructions. |
| EC05 | Performance | Massive Stack | `./push-swap "<500 random numbers>"` | Executes quickly without memory leaks. |
| EC06 | Stream Boundary | Checker with no instructions | `echo -n "" \| ./checker "1 2 3"` | Evaluates immediate state. Output: `OK\n`. |

---

## CLI, Core & Sorting Edge Cases

### 1. Input & Data Boundaries
These cases test the extremes of standard input parsing, ensuring the program does not index out of bounds or crash.

---

#### Edge Case EC01: Empty or Whitespace-Only String Argument
**Description:**  
Tests the parser's resilience when provided with an argument that exists but contains no valid integers (e.g., `""` or `"   "`). This is a classic trap: it is *not* the same as providing zero arguments.

**Input Argument:**
```bash
./push-swap ""
./push-swap "   "
./checker ""
./checker "   "
```
**Validation Rules:**
- **Distinction**: Running `./push-swap` (0 arguments) must display nothing. Running `./push-swap ""` or `./push-swap " "` provides an argument, but it cannot be parsed into valid integers.
- **Handling**: Because `""` or `" "` does not contain valid integers, it must output exactly `Error\n` to stderr and exit.
- **No Panic**: The parser must not panic (e.g., index out of bounds) when trying to split or parse empty/whitespace strings.

---

#### Edge Case EC02: Single Number Stack
**Description:**  
Tests the algorithm's base case when only a single element is provided.

**Input Argument:**
```bash
./push-swap "42"
```
**Validation Rules:**
- **Output Check:** Must print absolutely nothing, as a stack of 1 is already sorted.

---

#### Edge Case EC03: Already Sorted Stack
**Description:**  
Tests the algorithm's capability to evaluate the stack state before attempting any operations.

**Input Argument:**
```bash
./push-swap "0 1 2 3 4 5"
```
**Validation Rules:**
- **Output Check:** Must print absolutely nothing. No operations should be performed.

---

#### Edge Case EC04: Maximum and Minimum Integer Boundaries
**Description:**  
Tests the limits of the integer parser. Providing the maximum and minimum 32-bit (or 64-bit depending on system standard) integers must not trigger an overflow error if they are within standard Go `int` limits.

**Input Argument:**
```bash
./push-swap "2147483647 -2147483648 0"
```
**Validation Rules:**
- **Output Check:** Must output the correct sequence to sort the numbers without throwing a false `Error\n`.

---

### 2. Performance & Stream Boundaries

#### Edge Case EC05: Massive Stack Execution
**Description:**  
Tests the algorithm's performance, memory management, and time complexity. While the audit tests 100 numbers, the system should gracefully handle significantly larger inputs.

**Input Argument:**
```bash
./push-swap "<500 random different numbers>"
```
**Validation Rules:**
- **Execution Time:** Must complete the calculation within reasonable time limits (e.g., < 2 seconds).
- **Memory:** Must not leak memory or exhaust RAM during deep search trees.

---

#### Edge Case EC06: Checker Empty Instruction Stream
**Description:**  
Tests the checker's ability to evaluate the stack state when standard input immediately closes (EOF) with no instructions passed.

**Validation Rules:**
- **Already Sorted (`./checker "1 2 3"`):** Must output `OK\n`.
- **Unsorted (`./checker "3 2 1"`):** Must output `KO\n`.