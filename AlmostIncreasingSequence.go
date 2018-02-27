package main

func almostIncreasingSequence(sequence []int) bool {
	// Find first sequence violation and save two elems sequence in it
	skipIndxLeft := -1
	skipIndxRight := -1
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			skipIndxLeft = i
			skipIndxRight = i + 1
			break
		}
	}

	// Find out which element should be skipped
	skipIndx := -1
	if skipIndxLeft == 0 {
		skipIndx = skipIndxLeft
	} else if sequence[skipIndxLeft-1] < sequence[skipIndxRight] {
		skipIndx = skipIndxLeft
	} else {
		skipIndx = skipIndxRight
	}

	// Remove skipped element by some Golang magic
	sequence = append(sequence[:skipIndx], sequence[skipIndx+1:]...)

	// Find if there are other violation
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			return false
		}
	}
	return true
}
