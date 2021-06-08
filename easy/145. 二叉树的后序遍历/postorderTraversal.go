package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	f := func(node *TreeNode){}
	f = func(node *TreeNode){
		if node == nil {
			return
		}
		f(node.Left)
		f(node.Right)
		result = append(result,node.Val)
	}
	f(root)
	return result
}
