package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	if len(str) == 0 {
		return
	}
	l := len(str)
	count := 0
	ans := ""
	for i := 0; i < l; i++ {
		ans += string(str[i])
		count++
		if count == 8 {
			fmt.Printf("%s\n", ans)
			count = 0
			ans = ""
		}
	}
	if len(ans) != 0 {
		j := 8 - len(ans)
		for i := 0; i < j; i++ {
			ans += "0"
		}
		fmt.Printf("%s\n", ans)
	}
}
