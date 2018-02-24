package main

func checkPalindrome(inputString string) bool {
	if len(inputString) <= 1 {
		return true
	}

	for start, end := 0, len(inputString)-1; start < end; start, end = start+1, end-1 {
		if inputString[start] != inputString[end] {
			return false
		}
	}

	return true
}
