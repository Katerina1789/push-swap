# push-swap

![Go Version](https://img.shields.io/badge/Go-1.25.6-blue.svg)
![Standard Library](https://img.shields.io/badge/StdLib-Only-orange.svg)
![License](https://img.shields.io/badge/License-MIT-brightgreen.svg)
![Build Status](https://img.shields.io/badge/Build-InProcess-yellow.svg)
![Test Coverage](https://img.shields.io/badge/Coverage-0%25-darkred.svg)

[![01](https://img.shields.io/badge/zone01-Athens-916ADE?&labelColor=181717&style=for-the-badge&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBmaWxsPSJub25lIiBzdHJva2U9IndoaXRlIiBzdHJva2Utd2lkdGg9IjIiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIgc3Ryb2tlLWxpbmVqb2luPSJyb3VuZCI+PHBhdGggZD0iTTEyIDJMMiA3bDEwIDUgMTAtNS0xMC01eiIvPjxwYXRoIGQ9Ik0yIDE3bDEwIDUgMTAtNU0yIDEybDEwIDUgMTAtNSIvPjwvc3ZnPg==)](https://github.com/01-edu/public/tree/master/subjects/ascii-art-web)

A highly optimized sorting algorithm calculator and validation suite developed in Go. This project implements a solution to sort data on a stack, using a limited set of instructions, with the objective of achieving the lowest possible number of actions.

## Description

`push-swap` consists of two distinct programs:
1. **`push-swap`**: Calculates and displays the smallest sequence of instructions to sort a given stack of integers.
2. **`checker`**: Takes the initial integer stack as an argument and reads instructions from standard input. It executes the instructions on the stack and outputs `OK` if the result is sorted and the secondary stack is empty, or `KO` otherwise.

## Team & Roles

- **Systems Architect**: [Name] - Core Stack, Input Parsing, Memory Audit.
- **QA & Integration**: [Name] - Operations, Checker, Golden Test Pipeline.
- **Algorithm Specialist**: hmim - Sorting Heuristics, Benchmarking, Performance.

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

## Features

- **Optimized Sorting**: Specialized heuristics for small sets (2-5 numbers) and a scalable algorithm (e.g., Radix or Chunk Sort) for massive stacks.
- **Strict Validation**: Robust fail-fast parsing that rejects duplicates, non-integers, and overflows with a mandatory `Error` signal.
- **Shared Core Logic**: A unified internal engine ensuring that the `checker` and `push-swap` share the exact same mutation rules.
- **High Performance**: Pre-allocated memory structures to handle benchmarks of 100 and 500 integers without allocation overhead.

## Algorithm & Architecture

The project follows a **Modular CLI Pipeline** architecture, isolating core business logic within the `internal/` package:

- **`internal/stack`**: Manages state for stacks A and B using high-performance slice operations.
- **`internal/parser`**: Handles string splitting, whitespace sanitization, and strict integer conversion.
- **`internal/sorter`**: The algorithmic engine that determines the optimal path of instructions.
- **`internal/operations`**: Implements the 11 mandatory stack operations (`sa`, `sb`, `ss`, `pa`, `pb`, `ra`, `rb`, `rr`, `rra`, `rrb`, `rrr`).

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

Our testing suite follows Go idiomatic practices, utilizing package-local `testdata/` for comprehensive coverage:

- **Unit Tests**: Table-driven tests for all internal components.
- **Golden Tests**: Verification against mandatory audit scenarios and edge cases.
- **Fuzz Testing**: Stress-testing the parser with randomized input to ensure stability.
- **Benchmarks**: Tracking instruction counts and memory allocations.

To run the full suite:
```bash
go test -v ./...
```

## Performance Benchmarks

| Dataset Size | Instruction Limit | Expected Performance |
| :--- | :--- | :--- |
| 3 numbers | 3 | Sorted in ≤ 3 ops |
| 5 numbers | 12 | Sorted in ≤ 12 ops |
| 100 numbers | 700 | Sorted in < 700 ops |
| 500 numbers | 5500 | Sorted in < 5500 ops |

## Project Structure

```text
push-swap
├── cmd/
│   ├── checker/            # Entry point for the validation tool
│   │   └── main.go
│   └── push-swap/          # Entry point for the algorithm calculator
│       └── main.go
├── internal/               # Core business logic (Private packages)
│   ├── errs/               # Centralized error handling (Error\n)
│   ├── operations/         # Implementation of the 11 stack operations
│   ├── parser/             # Argument parsing, sanitization & validation
│   │   └── testdata/       # Fuzz seeds and input golden files
│   ├── sorter/             # Sorting algorithms (Dispatcher, Small, Big)
│   └── stack/              # High-performance Stack data structure
├── .ai/                    # AI collaboration logs & protocols
├── .docs/                  # Project requirements, test cases, and workflow
│   ├── .team/              # Team workflow and checklists
│   └── golden-tests.md     # Mandatory audit test cases
├── go.mod                  # Project dependencies (Standard Library only)
└── README.md
```
## Project Documentation References:

- [PRD](.docs/PRD.md)
- [Golden Tests](.docs/golden-tests.md)
- [*additional document related to the project specifications*]()
- [*additional document related to the project specifications*]()

---
*This project is part of the Zone01 Campus curriculum. It is built and maintained according to the guidelines specified in the `.docs/` directory.*