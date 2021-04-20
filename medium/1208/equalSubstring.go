package main

import "fmt"

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

func equalSubstring(s string, t string, maxCost int) int {
	left, right, sum, res := 0, 0, 0, 0
	//开始扩大窗口
	for right < len(s) {
		sum += abs(int(s[right]), int(t[right]))
		right++
		//不符合条件缩小窗口
		for sum > maxCost {
			sum -= abs(int(s[left]), int(t[left]))
			left++
		}
		res = max(res, right-left)
	}
	return res
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return -(a - b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
