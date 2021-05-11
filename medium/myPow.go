package medium


func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

//快速幂
func myPow1(x float64, n int) float64 {
	res := 1.0
	for i := n; i != 0; i /= 2 {
		if i%2 != 0 {
			res *= x
		}
		x *= x
	}
	if n < 0 {
		return 1.0 / res
	}
	return res
}
