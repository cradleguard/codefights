package main

func palindromeRearranging(inputString string) bool {
	m := make(map[rune]int)

	for _, v := range(inputString) {
		if _ , ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	meetIsOdd := false
	for _,v := range(m) {
		if v % 2 == 0 {
			continue
		} else {
			if len(inputString) % 2 == 0 {
				return false
			} else if meetIsOdd == true {
				return false
			} else {
				meetIsOdd = true
			}
		} 
	}

	return true
}

