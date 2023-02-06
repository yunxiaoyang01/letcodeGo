package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := map[string][]int{
		"A": {-1, 0},
		"S": {0, -1},
		"W": {0, 1},
		"D": {1, 0},
	}
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	arr := strings.Split(str, ";")
	x, y := 0, 0
	for _, v := range arr {
		if !valid(v) {
			continue
		}
		dir := string(v[0])
		step, err := strconv.Atoi(v[1:])
		if err != nil || step < 0 {
			continue
		}
		x += m[dir][0] * step
		y += m[dir][1] * step

	}
	fmt.Printf("%d,%d\n", x, y)
}
func valid(s string) bool {
	if len(s) > 3 {
		return false
	}
	if !(strings.HasPrefix(s, "A") || strings.HasPrefix(s, "S") || strings.HasPrefix(s, "W") || strings.HasPrefix(s, "D")) {
		return false
	}
	_, err := strconv.Atoi(s[1:])
	if err != nil {
		return false
	}
	return true
}
