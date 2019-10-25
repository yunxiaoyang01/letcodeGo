package easy

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		max_left := maxDepth(root.Left)
		max_right := maxDepth(root.Right)
		return int(math.Max(float64(max_left), float64(max_right))) + 1
	}

}
