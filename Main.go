package main

import "fmt"

func main() {
	input := "a(bcdefghijkl(mno)p)q"
	fmt.Println("Output: " + reverseParentheses(input))
	fmt.Println("Expected: " + "apmnolkjihgfedcbq")
}
