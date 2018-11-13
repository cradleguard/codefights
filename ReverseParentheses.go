package main

func reverseParentheses(s string) string {
	ret := make([]rune, 0)
	var openBracketsIndex stack
	for i := 0; i < len(s); i++ {
		v := rune(s[i])
		if v == '(' {
			openBracketsIndex.push(i)
		} else if v == ')' {
			startPosition := openBracketsIndex.pop()
			for k:= i - 1; k > startPosition; k-- {
				ret = append(ret, rune(s[k]))
			}
			i -= 2
			s =  s[:startPosition] + s[startPosition + 1 :i] + s[i + 1 : ]
		} else if len(openBracketsIndex) == 0 {
			ret = append(ret,v)
		}
	}
	return string(ret)
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
