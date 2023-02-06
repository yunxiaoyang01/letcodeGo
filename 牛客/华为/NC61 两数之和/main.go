package main

func main() {
}

func twoSum(numbers []int, target int) []int {
	// write code here
	temp := map[int]int{}
	for i, v := range numbers {
		if k, ok := temp[target-v]; ok {
			return []int{k + 1, i + 1}
		}
		temp[v] = i
	}
	return []int{}
}
