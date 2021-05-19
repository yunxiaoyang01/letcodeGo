package main

import "fmt"

func main() {
	arr := []int{2,3,1,6,7}
	fmt.Println(countTriplets(arr))

}

func countTriplets(arr []int) int {
	ans :=0
	n := len(arr)
	s := make([]int, n+1)
	for i, val := range arr {
		s[i+1] = s[i] ^ val
	}
	for i := 0; i < n; i++ {
		for k := i+1 ;k< n; k++ {
			if s[i] == s[k+1] {
				ans += k-i
			}
		}
	}
	return ans
}
