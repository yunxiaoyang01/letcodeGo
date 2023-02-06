package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s1 := input.Text()
	input.Scan()
	s2 := input.Text()
	ns1 := strings.ToLower(s1)
	ns2 := strings.ToLower(s2)
	count := 0
	for _, v := range ns1 {
		if string(v) == ns2 {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
