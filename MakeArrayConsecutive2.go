package main

import "sort"

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	ret := 0
	for i := 1; i <= len(statues)-1; i++ {
		if dif := statues[i] - statues[i-1]; dif > 1 {
			ret += dif - 1
		}
	}

	return ret
}
