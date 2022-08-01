package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1010"
	b := "1011"
	fmt.Printf("%s", addBinary(a, b))
}

// 题解 模拟 先取出a,b 中长度大的那个数字 开始循环，从末尾开始相加每位都是对2取余
// 然后对2 进行除等于
// 最后如果 carry ==1 证明最后一次相加时有进位，在补一个1
func addBinary(a string, b string) string {
	ans := ""
	lenA := len(a)
	lenB := len(b)
	n := max(lenA, lenB)
	carry := 0
	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
