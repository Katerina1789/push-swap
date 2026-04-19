# Team Workflow & Implementation Playbook: push-swap

Welcome to the **push-swap** team! This document is our single source of truth for how we collaborate, build, and deliver this project successfully. By adhering to these guidelines, we ensure high code quality, linear history, and predictable delivery.

## Project Overview & Role Assignment
Develop a highly optimized, non-comparative sorting algorithm calculator (`push-swap`) and an instruction validation program (`checker`) in Go using a limited set of stack operations.

### Team Delivery
- **[Name]**: Systems Architect (Core Stack, Input Parsing, Memory Audit)
- **[Name]**: QA & Integration Engineer (Operations, Checker, End-to-End Pipeline)
- **[Name]**: Algorithm Specialist (Small Sorts, Advanced Sorting, Benchmarking)

### Role Responsibilities
- **Systems Architect:** Focuses on Milestone 1 & 5. Responsible for building the robust input parser, argument sanitization (handling empty strings, duplicates, overflows), the core `Stack` data structures in `/internal`, and final memory/edge-case audits.
- **QA & Integration Engineer:** Focuses on Milestone 2 & 5. Responsible for implementing the exact logic for all 11 instructions (`pa`, `sa`, `rr`, etc.), building the `checker` executable that reads `stdin`, and running the End-to-End Golden Tests pipeline.
- **Algorithm Specialist:** Focuses on Milestones 3, 4 & 5. Responsible for the mathematical sorting logic. Builds the hardcoded small-sorts (3-5 numbers), the large-scale algorithm (e.g., Radix or Chunk Sort) for the 100/500 limits, and handles final algorithmic benchmarking.

## Project Architecture
We follow a strict KISS (Keep It Simple, Stupid) and YAGNI (You Aren't Gonna Need It) approach to our Go project layout:
- `/internal`: Contains the core business logic. 90% of our code lives here.
- `/cmd`: Contains minimal entry points (`main.go` for `push-swap` and `checker`).

## Technical Specifications

### Core Stack & Constraints
- **Environment:** Go 1.21+
- **Dependencies:** Standard library only. No external packages allowed.
- **Entry Point:** `main.go` inside `cmd/push-swap/` and `cmd/checker/` must act as simple orchestrators.

### Code Quality & Standards
- **Formatting:** Code MUST be formatted (e.g., with `gofmt`) before any commit.
- **Error Handling:** Use the centralized `internal/errs` package for all validation failures. Do not manually print to `stderr` to ensure 100% audit consistency. No `panic()` in production code.
- **Security & Performance:** Validate all inputs to prevent vulnerabilities. Write efficient code avoiding unnecessary allocations (e.g., pre-allocate slice capacities).
- **Documentation:** Public functions and types must have clear inline comments. Keep the `.docs/` directory updated when features change.

### Testing Requirements
- **Coverage:** All new features and bug fixes require corresponding unit tests.
- **Pre-commit Rule:** Tests must pass locally before pushing. Do not push failing code.
- **Test Data Management:** Use package-local `testdata/` directories to store complex inputs, golden files, and fuzz seeds. These directories are ignored by the Go build tool but are accessible to your tests.

#### Example Implementation
Example of a `testdata/` structure for the `internal/parser` package:

```text
internal/
    parser/
        parser.go
        validator.go
        parser_test.go
        validator_test.go
        testdata/
            fuzz/
                seed_corpus.txt  # Contains various valid and invalid input strings for fuzz testing.
            golden/
                small_input.golden  # Expected output for a small set of numbers.
                large_input.golden  # Expected output for a large set of numbers.
            bench/
                large_input.txt  # A file containing 500 unique random integers for benchmarking.
```

#### Example: Golden Test Using `testdata/`
```go
package parser

import (
    "os"
    "testing"
)

func TestParse_SmallInput(t *testing.T) {
    input := "3 1 2"
    expectedBytes, err := os.ReadFile("testdata/golden/small_input.golden")
    if err != nil {
        t.Fatalf("Failed to read golden file: %v", err)
    }
    expected := string(expectedBytes)
    // ... compare result with expected ...
}
```
*In this example, the test reads the expected output from a golden file. This keeps test code clean and allows you to easily update expected outputs without changing the test logic.*



## Workflow Implementation

### 1. Understanding the PRD, Edge Cases & Milestones
Before writing any implementation logic, thoroughly review `PRD.md`. You must also consult `edge-cases.md`, `error-cases.md`, `audit-cases.md`, and `golden-tests.md` to ensure your code accounts for required technical constraints and official audit scenarios.
  - **Read & Understand:** Review `PRD.md` to grasp core functional requirements, milestones, and architecture.
  - **Constraint Check:** Review `edge-cases.md`, `error-cases.md`, and `audit-cases.md` to ensure your approach handles known constraints and potential failures.
  - **Team Alignment:** Identify and clarify any ambiguities or newly discovered edge cases with the team before writing code.

### 2. Working with Tasks
1. Claim an unassigned Task Card in the `../../.tasks/` directory that fits your role.
2. Update the `STATUS` on the card to `IN PROGRESS`.
3. Update the card to `DONE` only when your work satisfies all *Acceptance Criteria* and the *Definition of Done*.
4. **Status Update Rule:** Update checklist items incrementally. Commit task status updates within the same logical commit as the related code or docs change.

The task files must always show the real project state so that any teammate can understand progress immediately.

### 3. Effective Go & Testing
Code quality is a collective responsibility. To maintain our standard of excellence, every teammate must adhere to the following testing protocols and idiomatic Go practices.
Before marking a task as **DONE** or pushing to the remote `dev` branch, you **must** verify your implementation against the official good practices checklists in `.docs/.team/checklists/`:  
- [CLI Good Practices Checklist](./checklists/CLI-Good-Practices-Checklist.md)
- [Testing Good Practices Checklist](./checklists/Testing-Good-Practices-Checklist.md)

*Recommended as well:*
- [Conventional Commits](./checklists/Conventional-Commits.md)
- [Git Workflow](./checklists/Git-Workflow.md)

**Idiomatic Table-Driven Testing**  
We follow the **Table-Driven Tests** pattern as the absolute standard for our Go logic.
- **Why:** It keeps our tests **DRY** (Don't Repeat Yourself), makes adding edge cases as simple as adding a new row to a slice, and provides clear, decoupled reporting via `t.Run()` sub-tests.
- **Implementation:** Always define a slice of anonymous structs containing `name`, `input`, and `expected` fields to iterate through your test logic.

**Essential Performance & Safety Tips**
- **Zero-Panic Policy:** Do not use `panic()` for flow control. Handle all errors gracefully. Use `t.Fatalf` in tests only when a failure makes further testing impossible; otherwise, prefer `t.Errorf`.
- **The Race Detector:** Always run your tests with the `-race` flag (e.g., `go test -race ./...`). We have a zero-tolerance policy for data races.
- **Performance Awareness:** For performance-critical logic, implement **Benchmarks** using `testing.B`. Use the `-benchmem` flag to monitor and minimize memory allocations (`allocs/op`).
- **Fuzzing for Robustness:** If your function handles complex string parsing, implement a `Fuzz` test to uncover edge cases.

### 4. Effective Git Workflow

- **Branching:** Switch to the branch `dev` for any feature implementation, keeping the `main` branch clean. Merge in `main` only after the project is completed and successfully tested by the whole team.

- **Conventional Commits:** Use short and descriptive commit messages. 
  - **Format:** `<type>[optional scope]: <description>`   
  - **Types:** `feat:`, `fix:`, `docs:`, `test:`, `chore:`, `refactor:`, `build:`
  - **Scopes:** `core`, `cli`, `checker`, `push-swap`, `parser`, `docs`, `tests`, `sorter`.
  - **Good Examples:**
    - `feat(parser): add core parsing logic for input arguments`
    - `fix(checker): reject unsupported invalid instructions`
    - `test(sorter): add benchmarks for 500 numbers`

- **Commit Rhythm:** Stage related work together instead of using one large mixed commit. Keep commits atomic and readable.

- **Push Discipline:** Push regularly after completing a meaningful checkpoint. Do not wait until the whole project is finished.

- **Conflict Resolution Strategy:** If your push is rejected due to remote changes, run `git pull --rebase origin dev`. Resolve any conflicts locally, test your code, and then push.

### 5. Code Review Process (Pre-Merge)
Since we do not use formal Pull Requests, Code Reviews must be conducted collaboratively before updating any change into the remote `dev` branch:
1. **Preparation:** The author ensures all tests pass (`go test ./...`), the code is formatted (`gofmt`), and the associated Task Card is fully updated.
2. **Peer Review:** The author invites at least one teammate to review the **code** and **good practices** application according to the provided **checklists**.
3. **Approval & Merge:** Once verbally or textually approved by the reviewer, the author pushes the `dev` branch to the remote.

### 6. Final Team Test (Pre-Release) (TASK-10)
Before merging the `dev` branch into `main` to finalize the project, the entire team must conduct a joint final test:
1. **Feature Freeze:** Halt all new development on the `dev` branch.
2. **Golden Test Run:** Run the entire `golden-tests.md` suite to ensure absolute output parity with expectations.
3. **Audit Simulation:** Have a teammate who did not write the core logic attempt to break the application using the `audit-cases.md`, `edge-cases.md`, and `error-cases.md` documents.
4. **Documentation Check:** Verify that the root `README.md` is complete, all AI logs are committed, and all `.tasks/` cards reflect a `DONE` status.
5. **Main Merge & Release:** Merge `dev` into `main` and tag the release (e.g., `v1.0.0`).

### 7. Personal AI Collaboration Protocol
Use AI as a project support tool, not as a replacement for engineering judgment. Every request must stay inside the scope of the school assignment.
   - **Individual Logging:** Each teammate MUST maintain their own personal AI log file inside `../../.ai/` (e.g., `.ai/hmim.ai.txt`). Do not rely on one shared log.
   - **Log Entry Content:** Must include Date, Model Used, Discussion Summary, Action Taken, and Affected Area/Task Card.
   - **Commit Rule:** If the AI work affects a task, update the relevant task file and commit your personal AI log together with the changes.

### 8. Team Reminder
- Keep the repository readable.
- Keep commits atomic.
- Keep AI usage transparent.
- Keep task files updated.
- Keep every change inside the assignment scope.
