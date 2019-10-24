package main

import "strings"

func main() {
	println(isIsomorphic("foo", "bar"))
}

func isIsomorphic(s string, t string) bool {
	//循环遍历字符串，看当前字符，在后面出现的位置如果两个字符串出现的位置不一样则false
	for i := 0; i < len(s); i++ {
		if strings.Index(s[i+1:], s[i:i+1]) != strings.Index(t[i+1:], t[i:i+1]) {
			return false
		}
	}
	return true
}
