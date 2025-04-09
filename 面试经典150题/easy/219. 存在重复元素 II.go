package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i,v := range nums {
		if j,ok := m[v];ok && abs(i,j)<=k {
			return true
		}
		m[v] = i
	}
	return false
}

func abs(a,b int) int {
	if a>b {
		return a-b
	}
	return b-a
}

// abs(i-j)<=k
// i-j <= k || j-i <=k
// i-k <=j || j< i+k
