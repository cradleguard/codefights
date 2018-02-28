package main

import (
	"sort"
)

func sortByHeight(a []int) []int {
	var people []int

	for _, v := range a {
		if v != -1 {
			people = append(people, v)
		}
	}
	sort.Ints(people)

	var result []int
	indxPeople := 0
	for _, v := range a {
		if v != -1 {
			result = append(result, people[indxPeople])
			indxPeople++
		} else {
			result = append(result, -1)
		}
	}

	return result
}
