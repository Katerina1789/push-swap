// Struct to represent the stack and state of the piles a & b.
// IMPORTANT NOTE: The Stack struct should be designed to efficiently handle the operations required by the push-swap algorithm, such as push, pop, swap, rotate, and reverse rotate.

// **Ensure the `Stack` struct or its constructor (`NewStack`) accepts an initial capacity.
// **Since you'll know the number of arguments from os.Args immediately,
// pre-allocating the slice capacity will prevent the overhead of slice growth during the 500-number sort,
// helping you meet the performance benchmarks.