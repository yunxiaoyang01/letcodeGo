package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	inOrder := func(node *TreeNode){}
	result := []int{}
	inOrder = func(node *TreeNode){
		if node == nil {
			return 
		}
		inOrder(node.Left)
		result = append(result, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return result[len(result)-k]
}
