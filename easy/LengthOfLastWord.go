package easy

import "strings"

//func main() {
//	println(lengthOfLastWord("a "))
//}

func lengthOfLastWord(s string) int {
	slen := 0
	s1 := strings.TrimSpace(s)
	sArray := strings.Split(s1, " ")
	length := len(sArray)
	if length != 0 {
		slen = len(sArray[length-1])
	}
	return slen
}
