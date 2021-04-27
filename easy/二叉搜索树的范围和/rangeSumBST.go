package main
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum :=0
	var inOrder func(*TreeNode)
	inOrder = func(tn *TreeNode){
		if tn !=nil{
			inOrder(tn.Left)
			if (tn.Val>=low&&tn.Val<=high){
				sum+=tn.Val
			}
			inOrder(tn.Right)
		}
	}
	inOrder(root)
	return sum
}