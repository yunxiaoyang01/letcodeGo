package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		num, _ := strconv.Atoi(str)
		if num == 0 {
			break
		}
		fmt.Printf("%d\n", fn(num))
	}
}

func fn(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fn(n-2) + 1
}
