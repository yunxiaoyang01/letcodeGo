package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	str := input.Text()
	ans, _ := strconv.Atoi(str)
	res := []int{}
	m := map[int]struct{}{}
	for i := 0; i < ans; i++ {
		input.Scan()
		str = input.Text()
		randNum, _ := strconv.Atoi(str)
		if _, ok := m[randNum]; !ok {
			res = append(res, randNum)
			m[randNum] = struct{}{}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	for _, v := range res {
		fmt.Printf("%d\n", v)
	}
}
