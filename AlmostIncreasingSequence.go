package main

func almostIncreasingSequence(sequence []int) bool {
	if len(sequence) == 2 {
		if sequence[1] >= sequence[0] {
			return true
		}

		return false
	}

	dublicate := make(map[int]int)
	dublicateCount := 0
	for i := 0; i <= len(sequence)-1; i++ {
		dublicate[sequence[i]]++
		if dublicate[sequence[i]] == 2 {
			dublicateCount++
		} else if dublicate[sequence[i]] == 3 {
			return false
		}
	}

	if dublicateCount > 1 {
		return false
	}

	candidatForDel := 0
	for i := 0; i <= len(sequence)-2; i++ {
		if sequence[i] <= sequence[i+1] {
			continue
		} else {
			candidatForDel++
		}
	}

	if candidatForDel > 1 {
		return false
	}

	for i := 1; i <= len(sequence)-3; i++ {
		if sequence[i] <= sequence[i+1] {
			continue
		} else {
			if sequence[i+1] < sequence[i-1] {
				return false
			}
		}
	}

	return true
}
