package main

import (
	"strconv"
	"fmt"
)

func main() {
	input := []int{1, 1, 1}
	fmt.Println("Output: " + strconv.Itoa(arrayChange(input)))
	fmt.Println("Expected: true")
}
