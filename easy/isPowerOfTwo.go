package easy

//func main() {
//	println(isPowerOfTwo(2))
//}
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n&(n-1) == 0 {
		return true
	}
	return false
}
