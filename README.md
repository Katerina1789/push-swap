# push-swap

[![Go Language](https://img.shields.io/badge/Go-1.25.7-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Standard Library](https://img.shields.io/badge/StdLib-Only-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://pkg.go.dev/std)
[![Build Status](https://img.shields.io/badge/Build-Passing-00C853?style=for-the-badge&logo=github&logoColor=white)](https://github.com)
[![Test Coverage](https://img.shields.io/badge/Coverage-71.7%25-FFA726?style=for-the-badge&logo=codecov&logoColor=white)](https://go.dev/blog/cover)
[![Zone01](https://img.shields.io/badge/Zone01-Athens-FF6B35?style=for-the-badge&logo=42&logoColor=white)](https://zone01.gr/)
[![License](https://img.shields.io/badge/license-MIT-green?style=for-the-badge&logo=opensourceinitiative&logoColor=white)](LICENSE)

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

This project requires **Go 1.25.7+**. To build the binaries, execute:

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
push-swap/
├── cmd/
│   ├── checker/                  # Entry point for the validation tool
│   │   └── main.go
│   └── push-swap/                # Entry point for the algorithm calculator
│       └── main.go
├── internal/                     # Core business logic (private packages)
│   ├── ops/                      # All 11 stack operations (sa, pb, ra, ...)
│   │   ├── ops.go
│   │   └── ops_test.go
│   ├── parser/                   # Argument parsing, sanitization & validation
│   │   ├── parser.go
│   │   ├── validator.go
│   │   ├── errs.go               # Centralized error handling (Error\n + os.Exit)
│   │   ├── parser_test.go
│   │   ├── validator_test.go
│   │   └── testdata/             # Test data for parser package
│   │       ├── fuzz/             # Fuzzing corpus and seeds
│   │       │   ├── FuzzParse/    # Fuzz test artifacts
│   │       │   └── seed_corpus.txt
│   │       └── golden/           # Golden files for parser tests
│   │           ├── parse_e01_error.golden
│   │           ├── parse_e02_error.golden
│   │           ├── parse_e03_error.golden
│   │           ├── parse_ec01_error.golden
│   │           └── parse_valid.golden
│   ├── sorter/                   # Sorting engine
│   │   ├── dispatcher.go         # Rank normalisation, routing, fast-path checks
│   │   ├── hardcoded.go          # BFS-optimal lookup tables for n=2..5
│   │   ├── chunk.go              # Butterfly sliding-window algorithm for n>5
│   │   └── sorter_test.go        # Full test suite (unit + benchmark)
│   └── stack/                    # Pre-allocated slice-based Stack
│       ├── stack.go
│       └── stack_test.go
├── .ai/                          # AI collaboration logs & protocols
│   ├── ebasou.ai.txt
│   ├── hmim.ai.txt
│   ├── kkasdana.ai.txt
│   └── README.md
├── .docs/                        # Project requirements, test cases, and workflow
│   ├── .team/                    # Team workflow and task checklists
│   │   ├── checklists/           # Good practices checklists
│   │   │   ├── CLI-Good-Practices-Checklist.md
│   │   │   ├── Conventional-Commits.md
│   │   │   ├── Git-Workflow.md
│   │   │   └── Testing-Good-Practices-Checklist.md
│   │   ├── tasks/                # Task tracking files (TASK-01 to TASK-10)
│   │   ├── README.md
│   │   └── TEAM_WORKFLOW.md
│   ├── audit-cases.md            # Official audit test cases
│   ├── edge-cases.md             # Edge case scenarios
│   ├── error-cases.md            # Error handling test cases
│   ├── golden-tests.md           # Comprehensive test suite
│   └── PRD.md                    # Product requirements document
├── .scripts/                     # Automation & quality control scripts
│   ├── check.sh                  # Comprehensive quality check pipeline
│   ├── install-tools.sh          # Install required Go tools
│   ├── Makefile                  # Build automation
│   └── README.md
├── .gitignore                    # Git ignore rules
├── coverage.out                  # Test coverage report (generated)
├── go.mod                        # Module declaration (Standard Library only)
├── LICENSE                       # MIT License
├── README.md                     # Project documentation (this file)
├── checker                       # Compiled checker binary (generated)
└── push-swap                     # Compiled push-swap binary (generated)
```
## Project Documentation

### Core Documentation
- **[Product Requirements Document (PRD)](.docs/PRD.md)** - Complete project specifications, architecture, and requirements
- **[Golden Tests](.docs/golden-tests.md)** - Comprehensive test suite with all audit, error, and edge cases
- **[Audit Cases](.docs/audit-cases.md)** - Official Zone01 audit test cases (C01-C15)
- **[Error Cases](.docs/error-cases.md)** - Error handling validation scenarios (E01-E05)
- **[Edge Cases](.docs/edge-cases.md)** - Boundary conditions and edge scenarios (EC01-EC06)

### Team Workflow
- **[Team Workflow](.docs/.team/TEAM_WORKFLOW.md)** - Development workflow, Git practices, and collaboration guidelines
- **[Task Tracking](.docs/.team/tasks/)** - Individual task cards (TASK-01 through TASK-10)
- **[Good Practices Checklists](.docs/.team/checklists/)** - CLI, Testing, Git, and Conventional Commits guidelines

### AI Collaboration
- **[AI Collaboration Guide](.ai/README.md)** - Guidelines for AI-assisted development
- **[Team AI Logs](.ai/)** - Individual AI collaboration logs (hmim.ai.txt, ebasou.ai.txt, kkasdana.ai.txt)

### Automation
- **[Scripts Documentation](.scripts/README.md)** - Quality control and automation tools

---
*This project is part of the Zone01 Campus curriculum.
It is built and maintained according to the guidelines specified in the `.docs/` directory.*