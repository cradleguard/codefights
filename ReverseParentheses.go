package main

func reverseParentheses(s string) string {
	var openBracketsIndex stack
	for i, v := range s {
		if v == '(' {
			openBracketsIndex.push(i)
		} else if v == ')' {

		}
	}
	return ""
}

type stack []int

func (s *stack) push(i int) {
	*s = append(*s, i)
}

func (s *stack) pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}
