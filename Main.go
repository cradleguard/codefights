package main

import (
	"strconv"
	"fmt"
)

func main() {
	input := "zaa"
	fmt.Println("Output: " + strconv.FormatBool(palindromeRearranging(input)))
	fmt.Println("Expected: true")
}
