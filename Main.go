package main

import (
	"fmt"
	"strings"
)

func main() {
	input := []string{"abc", "ded"}
	str := strings.Join(addBorder(input), ", ")
	fmt.Println("Output: [" + str + "]")
	fmt.Println("Expected: " + "[\"*****\", \"*abc*\",\"*ded*\",\"*****\"]")
}
