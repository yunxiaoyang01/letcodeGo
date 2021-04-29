package main

import (
	"fmt"
	"sort"
)


func main(){
	nums := []int{0,1,0,1,0,1,99}
	fmt.Println(singleNumber(nums))
}


func singleNumber(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	sort.Ints(nums)
	result :=0
	for i:=0;i<len(nums);i++{
		if i==0{
			if nums[i]!=nums[i+1]{
				result=nums[i]	
			}
		}else if (i==len(nums)-1){
			if nums[i]!=nums[i-1]{
				result = nums[i]
			}
		}else{
			if nums[i]!=nums[i-1]&&nums[i]!=nums[i+1]{
				result=nums[i]	
			}
		}
	}
	return result
}