package main

func main() {

}

func missingNumber(nums []int) int {
	for i, num := range nums {
		if i != num {
			return i
		}
	}
	return -1
}
