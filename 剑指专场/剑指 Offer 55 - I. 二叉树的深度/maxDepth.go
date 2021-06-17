package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 题解 
//	  递归遍历左节点的深度和右节点的深度，然后+1
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right))+1
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}
