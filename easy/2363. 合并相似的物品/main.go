package main

import (
	"fmt"
	"sort"
)

func main() {
	items1 := [][]int{{1, 1}, {4, 5}, {3, 8}}
	items2 := [][]int{{3, 1}, {1, 5}}
	fmt.Printf("%+v", mergeSimilarItems(items1, items2))
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	ansMap := map[int]int{}
	for _, item := range items1 {
		ansMap[item[0]] += item[1]
	}
	for _, item := range items2 {
		ansMap[item[0]] += item[1]
	}
	valArr := []int{}
	for k, _ := range ansMap {
		valArr = append(valArr, k)
	}
	sort.Ints(valArr)
	res := [][]int{}
	for _, v := range valArr {
		res = append(res, []int{v, ansMap[v]})
	}
	return res
}
