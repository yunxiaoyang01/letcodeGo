package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	n,m := len(matrix) ,len(matrix[0])
	result := make([][]int,m)
	for i,_ := range result{
		result[i] = make([]int,n)
		for j,_ := range result[i] {
			result[i][j] = -1
		}
	}

	for i,row := range matrix{
		for j,v := range row{
			result[j][i] = v
		}
	}
	return result
}

func main()  {
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	result := transpose(matrix)
	s :=fmt.Sprintf("%d",result)
	fmt.Println(s)
}