package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	f := func(node *TreeNode) {}
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		result = append(result, node.Val)
		f(node.Right)
	}
	f(root)
	return result
}
