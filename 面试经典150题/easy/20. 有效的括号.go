package main

import "fmt"

func main() {
	s := "([]))"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	k := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 && i <= len(s)-1 {
				return false
			}
			if stack[len(stack)-1] == k[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
