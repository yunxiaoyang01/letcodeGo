package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("%s", compressString("aabcccccaaa"))
}

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	if len(S) == 1 || len(S) == 2 {
		return S
	}
	index := S[0]
	count := 1
	str := strings.Builder{}
	for i := 1; i < len(S); i++ {
		if S[i] == index {
			count++
		} else {
			str.WriteByte(index)
			str.WriteString(strconv.Itoa(count))
			index = S[i]
			count = 1
		}
	}
	str.WriteByte(index)
	str.WriteString(strconv.Itoa(count))
	length := str.Len()
	if length < len(S) {
		return str.String()
	}
	return S
}
