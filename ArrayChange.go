package main

func arrayChange(inputArray []int) int {
	steps := 0
	for i,v := range(inputArray) {
		if i == 0 {
			continue
		}

		if v <= inputArray[i-1] {
			dif := abs(inputArray[i-1] - v)
			steps += dif + 1
			inputArray[i] += dif + 1
		}
	}
	return steps
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
