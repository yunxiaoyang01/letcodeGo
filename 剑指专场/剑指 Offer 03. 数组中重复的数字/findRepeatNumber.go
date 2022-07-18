package main

func findRepeatNumber(nums []int) int {
	s := map[int]int{}
	for _, num := range nums {
		s[num]++
	}
	for k, v := range s {
		if v > 1 {
			return k
		}
	}
	return -1
}
