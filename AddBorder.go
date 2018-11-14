package main

import (
	"strings"
)

func addBorder(picture []string) []string {
	height := len(picture)
	width := len(picture[0])
	
	ret := make([]string,0,height +2)
	ret = append(ret, strings.Repeat( "*", width +2 ))

	for _,v := range(picture) {
		ret = append(ret, "*" + v + "*")
	}

	ret = append(ret,strings.Repeat( "*", width +2 ))
	return ret
}
