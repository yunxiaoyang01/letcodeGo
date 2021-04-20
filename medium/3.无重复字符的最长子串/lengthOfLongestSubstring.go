package main

import "fmt"

func main(){
	fmt.Println(lengthOfLongestSubstring("abcbd"))
}

func lengthOfLongestSubstring(s string) int {
	n :=len(s)
	if n<1{
		return 0
	}
	right,ans:=0,1
	m :=map[byte]int{}
	for left := 0;left<n ;left++{
		if left!=0{
			delete(m,s[left-1])
		}
		for right < n && m[s[right]]==0{
			m[s[right]]=1
			right++
		}
		ans = max(ans,right-left)
	}
	return ans
}

func max(a,b int) int  {
	if a>b {
		return a
	}
	return b
}
