package main

func main() {
	
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var vals []int
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			vals = append(vals, root.Val)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	var vals1 []int
	vals1 = append(vals1, vals...)
	vals = []int{}
	dfs(root2)
	if len(vals) != len(vals1) {
		return false
	}
	for i, v := range vals {
		if v != vals1[i] {
			return false
		}
	}
	return true
}
