package main

import (
	"strconv"
)

func isLucky(n int) bool {
	var sum int64
	nstring := strconv.Itoa(n)
	middleIndex := (len(nstring) / 2) - 1
	for i, v := range nstring {
		dig, _ := strconv.ParseInt(string(v), 10, 32)
		if i <= middleIndex {
			sum += dig
		} else {
			sum -= dig
		}
	}

	return sum == 0
}
