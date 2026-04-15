# Golden Test Suite: push-swap

This document records the mandatory test cases used to verify the functional requirements, robust error handling, edge case resilience, and integration constraints of the `push-swap` and `checker` applications. It serves as the primary Definition of Done (DoD).

---

## Part 1: CLI & Core Golden Tests

### 1.1 Core Functionality & Audit Cases
These cases focus on the core functionality, standard usage patterns, and the official school audit rubric.

| Case ID | Description | Input Argument(s) | Expected Behavior / Output |
|:---:|---|---|---|
| C01 | No Arguments (push-swap) | `./push-swap` | Displays nothing. |
| C02 | Valid Sorting | `./push-swap "2 1 3 6 5 8"` | Valid solution & less than 9 instructions. |
| C03 | Already Sorted | `./push-swap "0 1 2 3 4 5"` | Displays nothing (0 instructions). |
| C04 | Invalid Input (String) | `./push-swap "0 one 2 3"` | `Error\n` on standard error. |
| C05 | Invalid Input (Duplicate) | `./push-swap "1 2 2 3"` | `Error\n` on standard error. |
| C06 | Benchmark: 5 Random Nums | `./push-swap "<5 random numbers>"` | Valid solution & less than 12 instructions. |
| C07 | No Arguments (checker) | `./checker` | Displays nothing. |
| C08 | Invalid Input (checker) | `./checker "0 one 2 3"` | `Error\n` on standard error. |
| C09 | Incorrect Instructions | `echo -e "sa\npb\nrrr\n" \| ./checker "0 9 1 8 2 7 3 6 4 5"` | `KO\n` on standard output. |
| C10 | Correct Instructions | `echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" \| ./checker "0 9 1 8 2"` | `OK\n` on standard output. |
| C11 | Integration (push-swap + checker) | `ARG="4 67 3 87 23"; ./push-swap "$ARG" \| ./checker "$ARG"` | `OK\n` on standard output. |
| C12 | Benchmark: 100 Random Nums | `ARG="<100 random numbers>"; ./push-swap "$ARG"` | Valid solution & less than 700 instructions. |
| C13 | Integration: 100 Random Nums | `ARG="<100 random numbers>"; ./push-swap "$ARG" \| ./checker "$ARG"` | `OK\n` on standard output. |

---

### 1.2 Error & Validation Cases
This phase validates robustness and strict enforcement of the input contract without crashing or exposing internal panics.

| Case ID | Program | Description | Input Argument(s) / Stdin | Expected Output |
|:---:|---|---|---|---|
| E01 | push-swap/checker | Non-integer argument | `./push-swap "0 one 2 3"` | `Error\n` on stderr |
| E02 | push-swap/checker | Duplicate integer | `./push-swap "1 2 2 3"` | `Error\n` on stderr |
| E03 | push-swap/checker | Integer overflow/underflow | `./push-swap "2147483648"` | `Error\n` on stderr |
| E04 | checker | Invalid instruction | `echo -e "sa\nbadcmd\n" \| ./checker "1 2 3"` | `Error\n` on stderr |
| E05 | checker | Incorrectly formatted instruction | `echo -e "sa \n" \| ./checker "1 2 3"` | `Error\n` on stderr |

---

### 1.3 Edge & Boundary Cases
These cases verify correct handling of valid edge cases, boundary conditions, and performance limits.

| Case ID | Category | Description | Input Argument(s) | Expected Behavior / Output |
|:---:|---|---|---|---|
| EC01 | Input Boundary | Empty or Whitespace String | `./push-swap ""` / `"   "` | `Error\n` on stderr. (Fails integer parsing). |
| EC02 | Input Boundary | Single Number | `./push-swap "42"` | Displays nothing (already sorted). |
| EC03 | Input Boundary | Already Sorted | `./push-swap "1 2 3 4"` | Displays nothing (0 instructions). |
| EC04 | Data Limits | Max/Min Int bounds | `./push-swap "2147483647 -2147483648 0"` | Parsed correctly and outputs valid sorting instructions. |
| EC05 | Performance | Massive Stack | `./push-swap "<500 random numbers>"` | Executes quickly without memory leaks. |
| EC06 | Stream Boundary | Checker with no instructions | `echo -n "" \| ./checker "1 2 3"` | Evaluates immediate state. Output: `OK\n`. |

---

### 1.4 Actionable Performance Benchmarks
These test cases provide the exact bash commands (using `shuf` to generate non-repeating random numbers) to verify the strict instruction count limits mentioned in the audit.

| Case ID | Category | Description | Command | Expected Output |
|:---:|---|---|---|---|
| P01 | Benchmark | 5 Random Nums (Limit: 12) | `ARG=$(shuf -i 1-100 -n 5 \| tr '\n' ' '); ./push-swap $ARG \| wc -l` | `< 12` |
| P02 | Benchmark | 100 Random Nums (Limit: 700) | `ARG=$(shuf -i 1-1000 -n 100 \| tr '\n' ' '); ./push-swap $ARG \| wc -l` | `< 700` |
| P03 | Benchmark | 500 Random Nums (Limit: 5500) | `ARG=$(shuf -i 1-10000 -n 500 \| tr '\n' ' '); ./push-swap $ARG \| wc -l` | `< 5500` |

---

## Part 2: Integration & Regression Verification

### 2.1 Verification Points

- **Data Races:** `go run -race .` should log zero data races during heavy logic execution.
- **Allowed Packages:** `go list -deps ./... | grep -v "^push-swap"` should show only standard library packages.
- **Coverage / Regression:** `go test ./... -cover` should achieve the minimum required code coverage.
- **Formatting / Static Checks:** `gofmt -l .` and `go vet ./...` produce no formatting issues or vet warnings.
- **Build Verification:** `go build -o push-swap ./cmd/push-swap` and `go build -o checker ./cmd/checker` result in clean compilations with no errors.

---

## Final Integration Checklist

- [ ] **Core Functionality:** Standard journeys and core logic produce expected valid sorting outputs or `OK`/`KO` statuses.
- [ ] **Error Handling:** Validation rules gracefully catch bad inputs, producing strict `Error\n` messages on `stderr` without exposing panics.
- [ ] **Edge Cases:** Boundary conditions (massive payloads, empty inputs, single strings) are sanitized and handled safely.
- [ ] **Integration Checks:** Package restrictions, test coverage, race detector, formatting, and build processes pass.
- [ ] **Deterministic Behavior:** All test cases produce consistent, repeatable results.
- [ ] **No Memory Leaks:** Program terminates cleanly without resource panics or memory leaks.