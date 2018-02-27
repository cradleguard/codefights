package main

func almostIncreasingSequence(sequence []int) bool {
	if len(sequence) == 2 {
		if sequence[1] >= sequence[0] {
			return true
		}

		return false
	}

	candidatForDel := make(map[int]int)
	for i := 0; i <= len(sequence)-2; i++ {
		if sequence[i] <= sequence[i+1] {
			continue
		} else {
			_, ok := candidatForDel[i]
			if !ok {
				candidatForDel[i] = sequence[i]
			}

		}
	}

	return len(candidatForDel) <= 1
}
