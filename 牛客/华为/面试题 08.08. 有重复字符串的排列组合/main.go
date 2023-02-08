package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%+v", permutation("qqe"))
}

func permutation(S string) []string {
	cs := []byte(S)
	sort.Slice(cs, func(i, j int) bool {
		return cs[i] < cs[j]
	})
	n := len(cs)
	res := []string{}
	visited := make([]bool, n)
	dfs(cs, "", visited, &res)
	return res
}
func dfs(cs []byte, path string, visited []bool, res *[]string) {
	if len(path) == len(cs) {
		*res = append(*res, path)
		return
	}
	for i := 0; i < len(cs); i++ {
		//如果arr[i] 已经加入当前排列，则不能多次加入当前排列。
		//如果当 i>0 时，arr[i]=arr[i−1] 且arr[i−1] 未加入当前排列，
		//则不能将arr[i] 加入当前排列，否则arr[i−1] 将在arr[i] 之后加入当前排列，导致出现重复排列。
		if visited[i] || (i > 0 && cs[i] == cs[i-1] && !visited[i-1]) {
			continue
		}
		visited[i] = true
		dfs(cs, path+string(cs[i]), visited, res)
		visited[i] = false
	}
}
