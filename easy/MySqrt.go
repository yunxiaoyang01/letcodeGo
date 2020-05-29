package easy

import "math"

//func main() {
//	println(mySqrt1(8))
//}
func mySqrt(x int) int {
	result := math.Sqrt(float64(x))
	k, _ := math.Modf(result)
	return int(k)
}

//二分法
func mySqrt1(x int) int {

	i, j := 1, x

	for i <= j {
		mid := (i + j) / 2
		if mid*mid > x {
			j = mid - 1
		} else if mid*mid < x {
			i = mid + 1
		} else {
			return mid
		}
	}
	return i - 1
}
