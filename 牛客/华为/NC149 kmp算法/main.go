package main

import "fmt"

func main() {
	S := "abab"
	T := "abacabab"
	fmt.Printf("%d\n", kmp(S, T))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算模板串S在文本串T中出现了多少次
 * @param S string字符串 模板串
 * @param T string字符串 文本串
 * @return int整型
 */
func kmp(S string, T string) int {
	ls := len(S)
	lt := len(T)
	if ls > lt {
		return 0
	}
	if ls == lt {
		if S == T {
			return 1
		} else {
			return 0
		}
	}
	arr := make([]string, 0)
	for i := 0; i < lt; i++ {
		if ls+i <= lt {
			arr = append(arr, T[i:ls+i])
		} else {
			break
		}
	}
	ans := 0
	for _, v := range arr {
		if v == S {
			ans++
		}
	}
	return ans
}
