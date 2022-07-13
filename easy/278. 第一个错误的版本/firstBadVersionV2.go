package main

func firstBadVersionV2(n int) int {
	if n == 1 {
		return 1
	}
	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
