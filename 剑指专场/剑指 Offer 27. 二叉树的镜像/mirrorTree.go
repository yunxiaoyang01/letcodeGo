package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
// 题解 ：
//		递归找到找到根部左右节点，然后将左右节点互换
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTreeNode := mirrorTree(root.Left)
	rightTreeNode := mirrorTree(root.Right)
	root.Left = rightTreeNode
	root.Right = leftTreeNode
	return root
}
