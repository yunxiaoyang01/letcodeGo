package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := input.Text()
		m := make(map[byte]int, 0)
		minCount := math.MaxInt
		for i := 0; i < len(str); i++ {
			m[str[i]]++
		}
		for _, v := range m {
			if v < minCount {
				minCount = v
			}
		}
		ans := &strings.Builder{}
		for i := 0; i < len(str); i++ {
			if m[str[i]] == minCount {
				continue
			}
			ans.WriteByte(str[i])
		}
		fmt.Printf("%s\n", ans.String())
	}
}
