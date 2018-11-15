package main

import (
	"strconv"
	"fmt"
)

func main() {
	//input := "zaa"
	yourLeft := 15
	yourRight := 10
	friendsLeft := 15
	friendsRight := 10
	fmt.Println("Output: " + strconv.FormatBool(areEquallyStrong(yourLeft,yourRight, friendsLeft, friendsRight)))
	fmt.Println("Expected: true")
}
