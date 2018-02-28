package main

import (
	"strings"
)

func commonCharacterCount(s1 string, s2 string) int {
	result := 0

	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				s2 = strings.Replace(s2, string(v1), "", 1)
				result++
				break
			}
		}
	}

	return result
}
