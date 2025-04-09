package main

func main() {
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	m1, n1 := m-1, n-1
	cur := m + n - 1
	for m1 >= 0 || n1 >= 0 {
		if m1 == -1 {
			nums1[cur] = nums2[n1]
			n1--
		} else if n1 == -1 {
			nums1[cur] = nums1[m1]
			m1--
		} else if nums1[m1] > nums2[n1] {
			nums1[cur] = nums1[m1]
			m1--
		} else {
			nums1[cur] = nums2[n1]
			n1--
		}
		cur--
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
