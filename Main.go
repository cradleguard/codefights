package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		[]int{0, 1, 1, 2},
		[]int{0, 5, 0, 0},
		[]int{2, 0, 3, 3},
	}
	fmt.Print(matrixElementsSum(input))
}
