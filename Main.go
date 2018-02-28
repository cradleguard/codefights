package main

import (
	"fmt"
)

func main() {
	input := []string{
		"aabcc",
		"adcaa",
	}
	fmt.Print(commonCharacterCount(input[0], input[1]))
}
