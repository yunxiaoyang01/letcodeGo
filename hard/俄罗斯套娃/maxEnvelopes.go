package main

import (
	"fmt"
	"sort"
)

func main(){
	envelopes := [][]int{{2,100},{3,200},{4,300},{5,500},{5,400},{5,250},{6,370},{6,360},{7,380}}
	fmt.Println(maxEnvelopes(envelopes))
} 

func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes,func(i,j int)bool{
       return  envelopes[i][0]<envelopes[j][0]
    })
    dp:=make([]int,len(envelopes))
	fmt.Printf("%d",envelopes)
   for i,_:=range dp{
       dp[i]=1
   }
    MM:=0
    for i:=0;i<len(dp);i++{
        
        for j:=0;j<i;j++{
            if envelopes[i][0]>envelopes[j][0]&&envelopes[i][1]>envelopes[j][1]{
                dp[i]=max(dp[i],dp[j]+1)
               
            }
        } 
        MM=max(MM,dp[i])
    }
    return MM
}
func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}