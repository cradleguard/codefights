package main

func reverseParentheses(s string) string {
	var openBracketsIndex stack
	for i := 0; i < len(s); i++ {
		v := rune(s[i])
		if v == '(' {
			openBracketsIndex.push(i)
		} else if v == ')' {
			startPosition := openBracketsIndex.pop()			
			s =  s[:startPosition] + reverse(s[startPosition + 1 : i]) + s[i + 1 : ]
			i -= 2
		} 
	}
	return s
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

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
