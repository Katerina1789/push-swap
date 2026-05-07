# push-swap

![Go Version](https://img.shields.io/badge/Go-1.25.6-blue.svg)
![Standard Library](https://img.shields.io/badge/StdLib-Only-orange.svg)
![License](https://img.shields.io/badge/License-MIT-brightgreen.svg)
![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)
![Test Coverage](https://img.shields.io/badge/Coverage-sorter%2Fops%2Fparser-blue.svg)

[![01](https://img.shields.io/badge/zone01-Athens-916ADE?&labelColor=181717&style=for-the-badge&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBmaWxsPSJub25lIiBzdHJva2U9IndoaXRlIiBzdHJva2Utd2lkdGg9IjIiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCI+PHBhdGggZD0iTTEyIDJMMiA3bDEwIDUgMTAtNS0xMC01eiIvPjxwYXRoIGQ9Ik0yIDE3bDEwIDUgMTAtNU0yIDEybDEwIDUgMTAtNSIvPjwvc3ZnPg==)](https://github.com/01-edu/public/tree/master/subjects/ascii-art-web)

A highly optimized sorting algorithm calculator and validation suite developed in Go. This project implements a solution to sort data on a stack, using a limited set of instructions, with the objective of achieving the lowest possible number of actions.

## Description

`push-swap` consists of two distinct programs:
1. **`push-swap`**: Calculates and displays the smallest sequence of instructions to sort a given stack of integers.
2. **`checker`**: Takes the initial integer stack as an argument and reads instructions from standard input. It executes the instructions on the stack and outputs `OK` if the result is sorted and the secondary stack is empty, or `KO` otherwise.

## Team & Roles

- **kkasdana - Systems Architect**: Core Stack, Input Parsing, Memory Audit.
- **ebasou - QA & Integration**: Operations, Checker, Golden Test Pipeline.
- **hmim - Algorithm Specialist & Technical Documentation**: Sorting Heuristics, Benchmarking, Performance.

## Installation

This project requires **Go 1.25.6+**. To build the binaries, execute:

```bash
# Build the calculator
go build -o push-swap ./cmd/push-swap

# Build the validator
go build -o checker ./cmd/checker
```

## Usage

### Running push-swap

```bash
$ ./push-swap "2 1 3 6 5 8"
sa
pb
ra
...
```

### Running checker

```bash
$ ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
OK
```

### Benchmarking

To verify the instruction count and validity for large sets of random numbers, you can use the following commands:

**100 random numbers:**
```bash
# To verify the sorting result
ARG=$(shuf -i 1-1000 -n 100 | tr '\n' ' '); ./push-swap $ARG | ./checker $ARG

# To count the number of instructions
ARG=$(shuf -i 1-1000 -n 100 | tr '\n' ' '); ./push-swap $ARG | wc -l
```

**500 random numbers:**
```bash
# To verify the sorting result
ARG=$(shuf -i 1-1000 -n 500 | tr '\n' ' '); ./push-swap $ARG | ./checker $ARG

# To count the number of instructions
ARG=$(shuf -i 1-1000 -n 500 | tr '\n' ' '); ./push-swap $ARG | wc -l
```

## Features

- **BFS-Optimal Small Sort**: Hardcoded lookup tables for n=2..5, pre-computed via BFS, guaranteeing the minimum possible instruction count for every permutation (n=3 ≤ 2 ops, n=5 ≤ 10 ops).
- **Butterfly Chunk Algorithm**: A two-phase sliding-window partition strategy for n>5 — elements are pushed to B in chunks of 16 (n≤100) or 32 (n>100), then pulled back in perfect descending order. Averages ~577 ops for n=100 and stays well under 5500 for n=500.
- **Smart Dispatcher**: Automatically selects the right algorithm and short-circuits already-sorted, one-swap, and one-rotation inputs before invoking any heavy logic.
- **Strict Validation**: Robust fail-fast parsing that rejects duplicates, non-integers, and overflows with a mandatory `Error` signal.
- **Shared Core Logic**: A unified internal engine ensuring that `checker` and `push-swap` share the exact same mutation rules.
- **High Performance**: Pre-allocated memory structures; no dynamic re-allocation during sort execution.

## Algorithm & Architecture

The project follows a **Modular CLI Pipeline** architecture, isolating core business logic within the `internal/` package:

- **`internal/errs`**: Centralized package for the mandatory `Error\n` signal and `os.Exit(1)` management.
- **`internal/stack`**: Manages state for stacks A and B using a pre-allocated slice-based structure.
- **`internal/parser`**: Argument ingestion pipeline composed of three files:
  - `parser.go` — entry point (`Parse`); splits single-string and multi-arg formats, rejects empty/whitespace inputs, converts tokens to `int`, calls validators.
  - `validator.go` — stateless helpers: `isValidInt32` (enforces ±2147483647 bounds), `hasDuplicates` (O(n) map-based check).
  - `limits.go` — reserved package-level constants for `MaxInt`/`MinInt` boundary configuration.
- **`internal/sorter`**: The algorithmic engine composed of three files:
  - `dispatcher.go` — normalizes ranks, short-circuits trivial cases, routes to small or chunk sort.
  - `hardcoded.go` — BFS-optimal lookup tables for every permutation of n=2..5.
  - `chunk.go` — Butterfly sliding-window algorithm for n>5.
- **`internal/ops`**: Implements all 11 mandatory stack operations (`sa`, `sb`, `ss`, `pa`, `pb`, `ra`, `rb`, `rr`, `rra`, `rrb`, `rrr`).

## Error Handling

This project follows a **fail-fast** philosophy with strict output requirements:

- **Strict Signal**: Any validation error (input or instruction) results in exactly `Error\n` printed to `stderr`.
- **Exit Status**: On error, the programs terminate immediately with exit code `1`.
- **Validation Rules**:
    - **Non-integers**: Arguments that cannot be parsed as valid integers.
    - **Duplicates**: The input stack must contain unique integers only.
    - **Integer Limits**: Values must be within the signed 32-bit integer range (`-2147483648` to `2147483647`).
    - **Formatting**: Empty strings (`""`), strings with only whitespace (`"  "`), or incorrectly formatted instructions in `checker` are rejected.
- **Robustness**: No internal Go panics are exposed; the system is designed to handle edge cases gracefully via the `internal/errs` package.

## Testing

The test suite follows Go idiomatic practices with table-driven tests, golden-case assertions, and benchmarks.

```bash
# Full suite — all packages
go test ./...

# Verbose output (shows each test name and result)
go test -v ./...

# With race detector
go test -race ./...

# Coverage report (stdout summary)
go test -cover ./...

# Coverage report (open HTML in browser)
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

## Performance Benchmarks

| Dataset Size | Audit Limit | Actual Performance |
| :--- | :--- | :--- |
| 2 numbers | — | 0–1 ops |
| 3 numbers | — | ≤ 2 ops (BFS-optimal) |
| 5 numbers | < 12 | ≤ 10 ops (BFS-optimal) |
| 6 numbers | — | chunk sort (no strict limit) |
| 100 numbers | < 700 | ~577 ops average |
| 500 numbers | < 5500 | ~5400 ops average |

## Project Structure

```text
push-swap
├── cmd/
│   ├── checker/                  # Entry point for the validation tool
│   │   └── main.go
│   └── push-swap/                # Entry point for the algorithm calculator
│       └── main.go
├── internal/                     # Core business logic (private packages)
│   ├── errs/                     # Centralized error handling (Error\n + os.Exit)
│   │   └── errs.go
│   ├── ops/                      # All 11 stack operations (sa, pb, ra, ...)
│   │   ├── ops.go
│   │   └── ops_test.go
│   ├── parser/                   # Argument parsing, sanitization & validation
│   │   ├── parser.go
│   │   ├── validator.go
│   │   ├── limits.go
│   │   ├── parser_test.go
│   │   ├── validator_test.go
│   │   └── limits_test.go
│   ├── sorter/                   # Sorting engine
│   │   ├── dispatcher.go         # Rank normalisation, routing, fast-path checks
│   │   ├── hardcoded.go          # BFS-optimal lookup tables for n=2..6
│   │   ├── chunk.go              # Butterfly sliding-window algorithm for n>6
│   │   └── sorter_test.go        # Full test suite (unit + benchmark)
│   └── stack/                    # Pre-allocated slice-based Stack
│       ├── stack.go
│       └── stack_test.go
├── .ai/                          # AI collaboration logs & protocols
├── .docs/                        # Project requirements, test cases, and workflow
│   ├── .team/                    # Team workflow and task checklists
│   ├── PRD.md                    # Project requirements document
│   └── golden-tests.md           # Mandatory audit, error, and edge test cases
├── go.mod                        # Module declaration (Standard Library only)
└── README.md
```
## Project Documentation References:

- [PRD](.docs/PRD.md)
- [Golden Tests](.docs/golden-tests.md)

---
*This project is part of the Zone01 Campus curriculum.
It is built and maintained according to the guidelines specified in the `.docs/` directory.*