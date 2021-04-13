package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main()  {
	nums :=[]int{3,30,34,5,9}
	fmt.Println(largestNumber(nums))
}


func largestNumber(nums []int) string {
	var numsStr []string
	for _,num := range nums{
		numsStr = append(numsStr,strconv.Itoa(num))
	}
	sort.Slice(numsStr, func(i, j int) bool {
		return numsStr[i]+numsStr[j]> numsStr[j]+numsStr[i]
	})
	if numsStr[0]=="0"{
		return "0";
	}
	return strings.Join(numsStr,"")
}