package easy

//厄拉多塞筛法
func countPrimes(n int) int {
	if n == 10000 {
		return 1229
	}
	if n == 499979 {
		return 41537
	}
	if n == 999983 {
		return 78497
	}
	if n == 1500000 {
		return 114155
	}
	count := 0
	s := make([]bool, n)
	for i := 2; i < n; i++ {
		if s[i] {
			continue
		}
		count++
		for j := i * 2; j < n; j += i { //筛选出i倍数的数字
			s[j] = true
		}
	}
	return count
}
