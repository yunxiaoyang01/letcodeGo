package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	target := input.Text()
	m := make(map[byte]struct{}, 0)
	for i := 0; i < len(target); i++ {
		m[target[i]] = struct{}{}
	}
	fmt.Printf("%d\n", len(m))
}
