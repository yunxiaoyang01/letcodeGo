package main

import "fmt"

func main() {
	matrix :=[][]int{{5,2},{1,6}}
	k := 1

	fmt.Println(kthLargestValue(matrix,k))
}

func kthLargestValue(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	results := make([]int, 0, m*n)
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	for i, row := range matrix {
		pre[i+1] = make([]int, n+1)
		for j, val := range row {
			pre[i+1][j+1] = pre[i+1][j] ^ pre[i][j+1] ^ pre[i][j] ^ val
			results = append(results,pre[i+1][j+1])
		}
	}
	return findKthLargestV2(results,k)
}


func findKthLargestV2(nums []int, k int) int {
	heapSize :=len(nums)
	buildMaxHeap(nums,heapSize)
	for i:=len(nums)-1;i>=len(nums)-k+1;i-- {
		nums[0],nums[i] = nums[i],nums[0]
		heapSize--
		maxHeap(nums,0,heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int,heapSize int)  {
	for i:=heapSize/2;i>=0;i-- {
		maxHeap(nums,i,heapSize)
	}
}

func maxHeap(nums []int,i ,heapSize int)  {
	l,r ,waitSwap := 2*i+1,2*i+2,i
	if l<heapSize && nums[l]> nums[waitSwap] {
		waitSwap = l
	}
	if r <heapSize && nums[r] > nums[waitSwap] {
		waitSwap = r
	}
	if waitSwap != i {
		nums[i],nums[waitSwap] = nums[waitSwap],nums[i]
		maxHeap(nums,waitSwap,heapSize)
	}
}