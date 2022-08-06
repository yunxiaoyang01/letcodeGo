package main

func main() {

}

func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for index, num := range nums {
		if 2*sum+num == total {
			return index
		}
		sum += num
	}
	return -1
}
