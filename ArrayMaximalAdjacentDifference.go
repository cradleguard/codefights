package main

func arrayMaximalAdjacentDifference(inputArray []int) int {
	maxDiff := 0
	for i :=1; i < len(inputArray); i++ {
		if abs(inputArray[i] - inputArray[i-1]) > maxDiff {
			maxDiff = abs(inputArray[i] - inputArray[i-1])
		}
	} 

	return maxDiff
}
