package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}
var res int
var pre int
func minDiffInBST(root *TreeNode) int {
	res = math.MaxInt64
	pre = 0
	midIter(root)
	return res
}
func midIter(root *TreeNode){
	if root==nil{
		return
	}
	midIter(root.Left)
	if pre !=0{
		res = min(res,root.Val-pre)
	}
	pre = root.Val
	midIter(root.Right)
}
func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}