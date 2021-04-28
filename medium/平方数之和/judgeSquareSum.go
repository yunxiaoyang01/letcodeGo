package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(judgeSquareSum(0))
}


func judgeSquareSum(c int) bool {
    for i:=0 ; i*i <=c;i++{
		b :=math.Sqrt(float64(c-i*i))
		if b == math.Floor(b){
			return true
		}
	}
	return false
}
