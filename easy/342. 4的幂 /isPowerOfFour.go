package main

func main() {

}

func isPowerOfFour(n int) bool {
	sum := 1
	for i := 0; i < 31; i++ {
		sum *= 4
		if sum == n {
			return true
		} else if sum > n {
			break
		}
	}
	return false
}
