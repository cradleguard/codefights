package main

import (
	"strconv"
	"fmt"
)

func main() {
	input := []int{50, 60, 60, 45, 70}
	res := alternatingSums(input)
	res0Str :=  strconv.Itoa(res[0])
	res1Str :=  strconv.Itoa(res[1])
	fmt.Println("Output: [" + res0Str + ", " + res1Str + "]")
	fmt.Println("Expected: " + "[180, 105]")
}
