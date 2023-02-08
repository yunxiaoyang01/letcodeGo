package main

func maxDepth(s string) int {
	ans := 0
	cur := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cur++
			if cur > ans {
				ans = cur
			}
		} else if s[i] == ')' {
			cur--
		}
	}
	return ans
}
