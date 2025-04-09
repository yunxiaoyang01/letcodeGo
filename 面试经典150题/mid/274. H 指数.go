package main

import (
	"fmt"
	"sort"
)

func main() {
	c := []int{1, 3, 1}
	fmt.Println(hIndex(c))
}

func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	i := len(citations) - 1
	for i >= 0 && h < citations[i] {
		h++
		i--
	}
	return h
}
