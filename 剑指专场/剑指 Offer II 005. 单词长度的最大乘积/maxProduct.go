package main

import "fmt"

func main() {
	words := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	fmt.Printf("%d", maxProduct2(words))
}

func maxProduct(words []string) int {
	var ans int

	for i := 0; i < len(words); i++ {
		word1 := words[i]
		for j := i + 1; j < len(words); j++ {
			word2 := words[j]
			if !hasSameChar(word1, word2) {
				ans = max(ans, len(word2)*len(word1))
			}
		}
	}
	return ans
}

func hasSameChar(word1 string, word2 string) bool {
	var bitMask1, bitMask2 = 0, 0
	for _, c := range word1 {
		bitMask1 |= (1 << (c - 'a'))
	}
	for _, c := range word2 {
		bitMask2 |= (1 << (c - 'a'))
	}
	return (bitMask1 & bitMask2) != 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 题解
// 预位运算+判断最大值
// 预先将数组中的每一个值的位运算值计算好然后存储在数组中
// 然后for循环两层，如果bitMask[i]和bitMask[j] 进行 & 如果等于0 证明两个元素之间没有相同字符
// 计算两个字符串的乘积与当前最大值进行比较，最终得到最大的乘积
func maxProduct2(words []string) int {
	var ans int
	n := len(words)
	bitMasks := make([]int, n)
	for i := 0; i < n; i++ {
		word := words[i]
		bitMask := 0
		for j := 0; j < len(word); j++ {
			bitMask |= (1 << (word[j] - 'a'))
		}
		bitMasks[i] = bitMask
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bitMasks[i]&bitMasks[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}
