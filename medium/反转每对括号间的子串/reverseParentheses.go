package main

import "fmt"

func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	stack := [][]byte{}
	str := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, str)
			str = []byte{}
		} else if s[i] == ')' {
			for j, n := 0, len(str); j < n/2; j++ {
				str[j], str[n-j-1] = str[n-j-1], str[j]
			}
			str = append(stack[len(stack)-1], str...)
			stack = stack[:len(stack)-1]
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
