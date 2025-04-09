package main

func main() {

}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index, num := range nums {
		k := target - num
		if v, ok := m[k]; ok {
			return []int{index, v}
		}
		m[num] = index
	}
	return []int{}
}
