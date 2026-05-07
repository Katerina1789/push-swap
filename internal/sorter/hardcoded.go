package sorter

import (
	"strconv"
	"strings"
)

// permKey returns a space-separated string of stack ranks top-to-bottom.
func permKey(a *Stack) string {
	parts := make([]string, a.Len())
	for i := 0; i < a.Len(); i++ {
		parts[i] = strconv.Itoa(a.At(i))
	}
	return strings.Join(parts, " ")
}

// sortSmall applies the BFS-optimal sort for n=2..6 (a already normalized).
func sortSmall(a, b *Stack, emit func(string)) {
	key := permKey(a)
	var instrs []string
	switch a.Len() {
	case 2:
		instrs = table2[key]
	case 3:
		instrs = table3[key]
	case 4:
		instrs = table4[key]
	case 5:
		instrs = table5[key]
	}
	for _, instr := range instrs {
		applyAndEmit(a, b, instr, emit)
	}
}

var table2 = map[string][]string{
	"0 1": {},
	"1 0": {"sa"},
}

var table3 = map[string][]string{
	"1 2 0": {"rra"},
	"2 0 1": {"ra"},
	"0 2 1": {"rra", "sa"},
	"2 1 0": {"ra", "sa"},
	"0 1 2": {},
	"1 0 2": {"sa"},
}

var table4 = map[string][]string{
	"0 1 2 3": {},
	"2 3 0 1": {"rra", "rra"},
	"1 3 0 2": {"sa", "ra", "sa"},
	"3 2 0 1": {"sa", "rra", "rra"},
	"2 1 0 3": {"sa", "ra", "sa", "rra", "sa"},
	"0 3 2 1": {"ra", "sa", "rra", "rra", "sa"},
	"1 0 3 2": {"rra", "sa", "ra", "sa", "rra", "sa"},
	"1 0 2 3": {"sa"},
	"1 2 3 0": {"rra"},
	"0 2 3 1": {"rra", "sa"},
	"3 1 0 2": {"ra", "sa"},
	"2 1 3 0": {"sa", "rra"},
	"2 0 3 1": {"sa", "rra", "sa"},
	"0 2 1 3": {"ra", "sa", "rra"},
	"3 1 2 0": {"rra", "sa", "ra"},
	"3 0 1 2": {"ra"},
	"2 3 1 0": {"rra", "rra", "sa"},
	"1 2 0 3": {"ra", "sa", "rra", "sa"},
	"3 2 1 0": {"sa", "rra", "rra", "sa"},
	"3 0 2 1": {"rra", "sa", "ra", "sa"},
	"2 0 1 3": {"sa", "ra", "sa", "rra"},
	"0 3 1 2": {"sa", "ra"},
	"1 3 2 0": {"ra", "sa", "rra", "rra"},
	"0 1 3 2": {"rra", "sa", "ra", "sa", "rra"},
}

var table5 = map[string][]string{
	"2 4 0 3 1": {"rra", "rra", "sa", "ra", "sa", "rra", "rra"},
	"1 2 0 3 4": {"ra", "sa", "rra", "sa"},
	"4 3 1 2 0": {"sa", "rra", "rra", "sa", "rra", "sa"},
	"0 4 3 2 1": {"ra", "sa", "ra", "ra", "sa", "ra", "sa", "rra", "sa"},
	"4 3 0 1 2": {"sa", "ra", "ra"},
	"2 3 4 1 0": {"rra", "rra", "sa"},
	"0 1 4 3 2": {"rra", "sa", "rra", "rra", "sa", "rra", "rra", "sa", "rra"},
	"2 3 4 0 1": {"rra", "rra"},
	"0 4 2 1 3": {"sa", "ra", "ra", "sa", "rra"},
	"0 1 3 2 4": {"ra", "ra", "sa", "rra", "rra"},
	"1 0 3 4 2": {"rra", "sa", "ra", "sa", "rra", "sa"},
	"0 4 1 3 2": {"sa", "rra", "rra", "sa", "rra", "rra"},
	"2 4 0 1 3": {"rra", "sa", "rra", "rra"},
	"0 1 3 4 2": {"rra", "sa", "ra", "sa", "rra"},
	"1 0 4 2 3": {"ra", "sa", "rra", "sa", "ra", "sa"},
	"3 1 4 2 0": {"ra", "sa", "ra", "ra", "sa", "rra", "sa"},
	"0 2 4 1 3": {"rra", "sa", "ra", "sa", "rra", "rra", "sa"},
	"4 3 2 0 1": {"sa", "ra", "ra", "sa", "ra", "sa", "rra"},
	"3 2 0 1 4": {"sa", "rra", "rra", "sa", "rra", "sa", "ra"},
	"0 1 2 3 4": {},
	"2 4 1 0 3": {"rra", "sa", "rra", "rra", "sa"},
	"1 4 2 0 3": {"sa", "ra", "ra", "sa", "rra", "sa"},
	"1 4 2 3 0": {"sa", "rra", "sa", "ra"},
	"4 0 1 3 2": {"rra", "rra", "sa", "rra", "rra"},
	"0 2 1 4 3": {"rra", "rra", "sa", "rra", "rra", "sa", "rra"},
	"3 0 2 1 4": {"rra", "sa", "rra", "rra", "sa", "rra"},
	"4 0 3 1 2": {"ra", "sa", "rra", "sa", "ra", "ra"},
	"3 1 2 0 4": {"rra", "sa", "rra", "rra", "sa", "rra", "sa"},
	"0 1 4 2 3": {"ra", "sa", "rra", "sa", "ra"},
	"2 1 3 4 0": {"sa", "rra"},
	"4 1 2 3 0": {"rra", "sa", "ra"},
	"3 0 1 2 4": {"rra", "sa", "ra", "ra"},
	"4 2 1 0 3": {"ra", "sa", "ra", "sa", "rra", "sa"},
	"1 4 3 0 2": {"ra", "sa", "rra", "rra", "sa", "rra"},
	"1 3 2 0 4": {"ra", "sa", "ra", "ra", "sa", "ra"},
	"2 1 4 3 0": {"sa", "ra", "ra", "sa", "ra", "ra"},
	"3 4 0 1 2": {"ra", "ra"},
	"0 2 1 3 4": {"ra", "sa", "rra"},
	"3 2 1 0 4": {"sa", "rra", "rra", "sa", "rra", "sa", "ra", "sa"},
	"2 4 1 3 0": {"sa", "rra", "sa", "ra", "ra", "sa", "rra"},
	"4 0 1 2 3": {"ra"},
	"0 4 1 2 3": {"sa", "ra"},
	"3 2 4 0 1": {"sa", "rra", "rra"},
	"0 3 4 1 2": {"rra", "sa", "rra", "sa"},
	"4 0 2 1 3": {"ra", "ra", "sa", "rra"},
	"1 0 2 4 3": {"rra", "rra", "sa", "ra", "ra", "sa"},
	"3 4 2 1 0": {"ra", "ra", "sa", "ra", "sa", "rra", "sa"},
	"4 3 2 1 0": {"sa", "ra", "ra", "sa", "ra", "sa", "rra", "sa"},
	"0 1 2 4 3": {"rra", "rra", "sa", "ra", "ra"},
	"3 2 1 4 0": {"sa", "ra", "ra", "sa", "ra", "sa"},
	"1 3 4 2 0": {"rra", "rra", "sa", "ra", "sa", "rra"},
	"4 1 3 2 0": {"rra", "sa", "rra", "rra", "sa", "rra", "rra"},
	"2 0 1 4 3": {"sa", "rra", "rra", "sa", "rra", "rra", "sa", "rra"},
	"2 0 3 4 1": {"sa", "rra", "sa"},
	"3 2 4 1 0": {"sa", "rra", "rra", "sa"},
	"0 3 2 4 1": {"ra", "sa", "rra", "rra", "sa"},
	"0 4 2 3 1": {"sa", "rra", "sa", "ra", "sa"},
	"4 2 0 1 3": {"ra", "sa", "ra", "sa", "rra"},
	"1 2 0 4 3": {"rra", "rra", "sa", "rra", "rra", "sa", "rra", "sa"},
	"2 4 3 1 0": {"ra", "sa", "ra", "ra", "sa"},
	"3 2 0 4 1": {"sa", "ra", "ra", "sa", "ra"},
	"1 0 3 2 4": {"ra", "ra", "sa", "rra", "rra", "sa"},
	"4 2 1 3 0": {"rra", "sa", "ra", "ra", "sa", "rra"},
	"0 4 3 1 2": {"ra", "sa", "rra", "rra", "sa", "rra", "sa"},
	"1 4 0 2 3": {"sa", "ra", "sa"},
	"2 4 3 0 1": {"ra", "sa", "ra", "ra"},
	"3 0 1 4 2": {"ra", "ra", "sa", "rra", "sa", "ra"},
	"2 1 0 3 4": {"sa", "ra", "sa", "rra", "sa"},
	"4 1 2 0 3": {"ra", "ra", "sa", "rra", "sa"},
	"4 3 0 2 1": {"sa", "rra", "rra", "sa", "rra"},
	"0 3 1 2 4": {"sa", "rra", "sa", "ra", "ra"},
	"1 3 2 4 0": {"ra", "sa", "rra", "rra"},
	"0 2 3 1 4": {"rra", "rra", "sa", "ra", "sa"},
	"4 1 0 3 2": {"rra", "rra", "sa", "rra", "rra", "sa"},
	"1 2 4 0 3": {"rra", "sa", "ra", "sa", "rra", "rra"},
	"4 0 3 2 1": {"rra", "sa", "rra", "rra", "sa", "rra", "rra", "sa"},
	"3 1 4 0 2": {"ra", "sa", "ra", "sa"},
	"4 2 3 0 1": {"rra", "sa", "rra", "sa", "ra"},
	"0 3 1 4 2": {"rra", "sa", "rra", "rra", "sa", "ra", "sa"},
	"3 4 1 0 2": {"ra", "ra", "sa"},
	"1 0 2 3 4": {"sa"},
	"4 0 2 3 1": {"rra", "sa", "ra", "sa"},
	"2 3 0 4 1": {"ra", "ra", "sa", "ra"},
	"3 4 1 2 0": {"rra", "rra", "sa", "rra", "sa"},
	"3 1 2 4 0": {"sa", "ra", "sa", "rra", "rra"},
	"1 3 0 4 2": {"rra", "sa", "rra", "rra", "sa", "ra"},
	"0 2 3 4 1": {"rra", "sa"},
	"1 2 3 0 4": {"rra", "rra", "sa", "ra"},
	"0 3 4 2 1": {"rra", "rra", "sa", "ra", "sa", "rra", "sa"},
	"2 0 4 3 1": {"sa", "ra", "ra", "sa", "ra", "ra", "sa"},
	"1 0 4 3 2": {"rra", "sa", "rra", "rra", "sa", "rra", "rra", "sa", "rra", "sa"},
	"3 4 0 2 1": {"rra", "rra", "sa", "rra"},
	"0 2 4 3 1": {"ra", "ra", "sa", "ra", "ra", "sa"},
	"3 4 2 0 1": {"ra", "ra", "sa", "ra", "sa", "rra"},
	"3 0 4 2 1": {"ra", "sa", "ra", "ra", "sa", "rra"},
	"4 2 3 1 0": {"rra", "sa", "rra", "sa", "ra", "sa"},
	"2 1 0 4 3": {"sa", "rra", "rra", "sa", "rra", "rra", "sa", "rra", "sa"},
	"1 3 4 0 2": {"rra", "sa", "rra"},
	"4 3 1 0 2": {"sa", "ra", "ra", "sa"},
	"3 1 0 2 4": {"rra", "sa", "ra", "ra", "sa"},
	"1 2 4 3 0": {"ra", "ra", "sa", "ra", "ra"},
	"3 0 4 1 2": {"ra", "sa", "ra"},
	"2 3 0 1 4": {"rra", "rra", "sa", "rra", "sa", "ra"},
	"2 3 1 0 4": {"rra", "rra", "sa", "rra", "sa", "ra", "sa"},
	"2 3 1 4 0": {"ra", "ra", "sa", "ra", "sa"},
	"4 2 0 3 1": {"rra", "sa", "ra", "ra", "sa", "rra", "sa"},
	"1 4 3 2 0": {"ra", "sa", "ra", "ra", "sa", "ra", "sa", "rra"},
	"2 0 1 3 4": {"sa", "ra", "sa", "rra"},
	"2 1 3 0 4": {"sa", "rra", "rra", "sa", "ra"},
	"2 0 3 1 4": {"sa", "rra", "rra", "sa", "ra", "sa"},
	"1 3 0 2 4": {"sa", "rra", "sa", "ra", "ra", "sa"},
	"2 0 4 1 3": {"rra", "sa", "ra", "ra", "sa", "ra"},
	"3 1 0 4 2": {"ra", "ra", "sa", "rra", "sa", "ra", "sa"},
	"2 1 4 0 3": {"rra", "sa", "ra", "ra", "sa", "ra", "sa"},
	"0 3 2 1 4": {"ra", "sa", "ra", "ra", "sa", "ra", "sa"},
	"1 2 3 4 0": {"rra"},
	"4 1 0 2 3": {"ra", "sa"},
	"3 0 2 4 1": {"sa", "ra", "sa", "rra", "rra", "sa"},
	"1 4 0 3 2": {"sa", "rra", "rra", "sa", "rra", "rra", "sa"},
	"4 1 3 0 2": {"ra", "sa", "rra", "sa", "ra", "ra", "sa"},
}
