package main

func allLongestStrings(inputArray []string) []string {
	maxlength := 0
	ret := make([]string, 0)
	for _, v := range inputArray {
		if len(v) < maxlength {
			continue

		} else if len(v) > maxlength {
			ret = make([]string, 0)
			maxlength = len(v)
		}
		ret = append(ret, v)
	}
	return ret
}
