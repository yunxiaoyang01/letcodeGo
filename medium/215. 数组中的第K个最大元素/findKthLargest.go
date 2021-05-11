package main

import (
	"fmt"
	"sort"
)

func main()  {
	nums :=[]int{3,2,3,1,2,4,5,5,6}
	k:=4
	fmt.Println(findKthLargestV2(nums,k))
}

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]>nums[j]
	})
	return nums[k-1]
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

