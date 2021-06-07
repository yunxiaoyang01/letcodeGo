package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	f := func(node *TreeNode) {}
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		f(node.Left)
		f(node.Right)
	}
	f(root)
	return result
}
