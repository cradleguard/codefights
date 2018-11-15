package main

import (
	"strconv"
	"fmt"
)

func main() {
	input := []int{-1, 4, 10, 3, -2}
	fmt.Println("Output: " + strconv.Itoa(arrayMaximalAdjacentDifference(input)))
	fmt.Println("Expected: 7")
}
