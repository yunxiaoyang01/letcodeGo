package main

import "fmt"

func main(){
	arr :=[]int{4,8,2,10}
	queries := [][]int{{2,3},{1,3},{0,0},{0,3}}
	fmt.Printf("%d",xorQueries(arr,queries))
	fmt.Printf("%d",xorQueriesV2(arr,queries))
}

func xorQueries(arr []int, queries [][]int) []int {
	result:=[]int{}
    for _,query := range queries {
		if query[0]==query[1]{
			result = append(result, arr[query[0]])
		}else{
			subResult := 0
			for i:=query[0];i<=query[1];i++{
				subResult^=arr[i]
			}
			result =append(result, subResult)
		}
	}
	return result
}
// 最优解，利用前缀和，先算出来0到每个下标下的异或结果，存储到数组中。
// 当 left=0 时，Q(left,right)=xors[right+1]
// 当 left!=0 时，Q(left,right) = xors[left]^xors[right+1]
// 根据异或运算的结合律 ，当left = 0 是 也满足Q(left,right) = xors[left]^xors[right+1]
//根据上述分析，这道题可以分两步求解。

// 计算前缀异或数组 xors；

// 计算每个查询的结果，第 i 查询的结果为xors[queries[i][0]]^xors[queries[i][1]+1]
func xorQueriesV2(arr []int, queries [][]int) []int {
    xors := make([]int, len(arr)+1)
    for i, v := range arr {
        xors[i+1] = xors[i] ^ v
    }
    ans := make([]int, len(queries))
    for i, q := range queries {
        ans[i] = xors[q[0]] ^ xors[q[1]+1]
    }
    return ans
}