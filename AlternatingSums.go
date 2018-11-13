package main

func alternatingSums(a []int) []int {
	ret := []int{0,0}
	for i, v := range(a) {
		if i % 2 == 0 {
			ret[0] = ret[0] + v
		} else {
			ret[1] = ret[1] + v
		}
	} 

	return ret
}