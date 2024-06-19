package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "               "
	l := 5
	fmt.Printf("%s", replaceSpaces(S, l))
}

func replaceSpaces(S string, length int) string {
	if len(S)> length {
		S = S[0:length]
	}else {
		S = buchong(S,length)
	}
	ans := strings.Replace(S, " ", "%20", -1)
	return ans
}

func buchong(S string, length int) string{
	sub := length-len(S)
	sw := strings.Builder{}
	for i:=0;i<sub;i++{
		sw.WriteString(" ")
	}
	return S+sw.String()
}
