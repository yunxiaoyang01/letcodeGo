package main

func main() {

}

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	var res int
	for key, value := range m {
		if value == 1 {
			res = key
		}
	}
	return res

}
