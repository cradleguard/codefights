package main

func areSimilar(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	swapped := false
	for i, v := range a {
		if v != b[i] {
			if swapped == true {
				return false
			}

			for k := i + 1; k < len(b); k++ {
				if v == b[k] && a[k] != b[k] {
					b[k] = b[i]
					b[i] = v
					swapped = true
					break
				}
			}

			if swapped == false {
				return false
			}
		}
	}

	return true
}
