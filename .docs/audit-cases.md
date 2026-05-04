# Audit Cases: push-swap

This document outlines the standard audit cases to verify the core functional requirements, benchmark outcomes, and execution behavior of the `push-swap` and `checker` programs based on the official project audit.

---

### Summary Table of Audit Cases

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
| C12 | Benchmark: 100 Random Nums | `ARG="<100 random numbers>"; ./push-swap "$ARG"` | Valid solution & less than `700` instructions. |
| C13 | Integration: 100 Random Nums | `ARG="<100 random numbers>"; ./push-swap "$ARG" \| ./checker "$ARG"` | `OK\n` on standard output. |
| C14 | Benchmark: 500 Random Nums | `ARG="<500 random numbers>"; ./push-swap "$ARG"` | Valid solution & less than `5500` instructions. |
| C15 | Integration: 100 Random Nums | `ARG="<500 random numbers>"; ./push-swap "$ARG" \| ./checker "$ARG"` | `OK\n` on standard output. |

---

## CLI & Core Audit Cases Detailed

### Audit Case C02: Valid Sorting
**Description:**  
Tests the `push-swap` program's ability to generate an efficient set of instructions for a known unsorted list.

**Input Argument:**
```bash
./push-swap "2 1 3 6 5 8"
```
**Validation Rules:**
- **Output Check:** Must output a valid sequence of instructions separated by `\n`.
- **Instruction Count:** The output must contain fewer than 9 instructions.

---

### Audit Case C04 & C05: Error Handling
**Description:**  
Tests the program's strict adherence to input validation (rejecting non-integers and duplicate numbers).

**Input Argument:**
```bash
./push-swap "0 one 2 3"
./push-swap "1 2 2 3"
```
**Validation Rules:**
- **Output Check:** Must print exactly `Error` followed by a newline (`\n`).
- **Output Stream:** The error must be printed to the standard error (`stderr`), not standard output (`stdout`).

---

### Audit Case C06: Benchmark (5 Numbers)
**Description:**  
Tests the `push-swap` algorithm's performance against the strict audit limit for 5 random numbers.

**Input Argument:**
```bash
./push-swap "<5 random different numbers>"
```
**Validation Rules:**
- **Instruction Count:** Must return a valid sorting sequence with less than 12 instructions.

---

### Audit Case C09 & C10: Checker Instruction Validation
**Description:**  
Tests the `checker` program's ability to read instructions from standard input, apply them, and verify the final state of stack `a`.

**Input Argument (KO scenario):**
```bash
echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
```
**Input Argument (OK scenario):**
```bash
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
```
**Validation Rules:**
- **Output Format:** Must output exactly `KO\n` or `OK\n` to standard output.

---

### Audit Case C11, C13 & C15: Full Integration
**Description:**  
Tests the seamless pipeline between both programs. `push-swap` generates instructions which are immediately piped to `checker`.

**Input Argument:**
```bash
ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"
```
**Validation Rules:**
- **Output:** The final output from the checker must be `OK\n`.

---

### Audit Case C12: Benchmark (100 Numbers)
**Description:**  
Tests the algorithm's performance and efficiency at a medium scale.

**Input Argument:**
```bash
ARG="<100 random different numbers>"; ./push-swap "$ARG" | wc -l
```
**Validation Rules:**
- **Instruction Count:** The total number of instructions outputted by `push-swap` must be strictly less than 650.
### Audit Case C12: Benchmark (100 Numbers)
**Description:**  
Tests the algorithm's performance and efficiency at a medium scale.

**Input Argument:**
```bash
ARG=$(shuf -i 1-1000 -n 100 | tr '\n' ' '); ./push-swap $ARG | wc -l
```
**Validation Rules:**
- **Instruction Count:** The total number of instructions outputted by `push-swap` must be strictly less than 650.

### Audit Case C14: Benchmark (500 Numbers)
**Description:**  
Tests the algorithm's performance and efficiency at a medium scale.

**Input Argument:**
```bash
ARG=$(shuf -i 1-1000 -n 500 | tr '\n' ' '); ./push-swap $ARG | wc -l
```
**Validation Rules:**
- **Instruction Count:** The total number of instructions outputted by `push-swap` must be strictly less than 5500.