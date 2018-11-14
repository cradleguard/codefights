package main

import (
	"strconv"
	"fmt"
)

func main() {
	inputA := []int{2, 1, 2, 1, 1, 1, 2}
	inputB := []int{1, 1, 2, 1, 2, 1, 2}
	fmt.Println("Output: " + strconv.FormatBool(areSimilar(inputA,inputB)))
	fmt.Println("Expected: true")
}
