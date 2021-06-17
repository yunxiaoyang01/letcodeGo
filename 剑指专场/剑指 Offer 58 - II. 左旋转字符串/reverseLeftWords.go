package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s, k))
}

// 切片解法
func reverseLeftWordsV1(s string, n int) string {
	s0 := s[0:n]
	s1 := s[n:]
	return s1 + s0
}

// 暴力拼接
func reverseLeftWordsV2(s string, n int) string {
	sArray := []rune(s)
	res := []rune{}
	for i := n; i < len(sArray); i++ {
		res = append(res, sArray[i])
	}

	for i := 0; i < n; i++ {
		res = append(res,sArray[i])
	}
	return string(res)
}



// 求余算法 简化暴力拼接
func reverseLeftWords(s string, n int) string {
	sArray := []rune(s)
	res := []rune{}
	for i := n; i < n+len(sArray); i++ {
		res = append(res, sArray[i%len(sArray)])
	}
	return string(res)
}
