package main

import (
	"fmt"
)

func main() {
	// batch values into groups of size `size`
	var (
		vs      = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}     // slice
		size    = 3                                       // batch size
		batches = make([][]int, 0, (len(vs)+size-1)/size) // allocate memory for batches
	)
	sliceInfo("vs        ", vs)
	for size < len(vs) {
		vs, batches = vs[size:], append(batches, vs[0:size:size])
	}
	batches = append(batches, vs)

	// even if we empty vs, we still reference the underlying array in batches
	vs = nil

	sliceInfo("vs        ", vs)
	sliceInfo("batches   ", batches)
	sliceInfo("batches[0]", batches[0])
	sliceInfo("batches[1]", batches[1])
}

// [vs        ] 0xc0000be000 len=10 cap=10 [0 1 2 3 4 5 6 7 8 9]
// [vs        ] 0xc0000be048 len=1 cap=1 [9]
// [batches   ] 0xc0000ba120 len=4 cap=4 [[0 1 2] [3 4 5] [6 7 8] [9]]
// [batches[0]] 0xc0000be000 len=3 cap=3 [0 1 2]
// [batches[1]] 0xc0000be018 len=3 cap=3 [3 4 5]

func sliceInfo(name string, s interface{}) {
	switch v := s.(type) {
	case [][]int:
		fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, v, len(v), cap(v), v)
	case []int:
		fmt.Printf("[%s] %p len=%d cap=%d %v\n", name, v, len(v), cap(v), v)
	}
}
