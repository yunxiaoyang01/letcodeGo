package easy

/**
	照镜子原理
	左          右
	1          1
   / \        / \
  2   2      2   2
 / \ / \    / \ / \
3  4 4  3  3  4 4  3

	true
	左          右
 	1			1
   / \		   / \
  2   2       2   2
   \   \	   \   \
   3    3       3   3
	  false

*/
func isSymmetric(root *TreeNode) bool {
	return isM(root, root)
}
func isM(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && isM(a.Left, b.Right) && isM(a.Right, b.Left)
}
