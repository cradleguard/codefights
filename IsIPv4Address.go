package main

import (
	"strings"
)

func isIPv4Address(inputString string) bool {
	splitted := strings.Split(inputString, ".")

	if len(splitted) != 4 {
		return false
	}

	return true
}
