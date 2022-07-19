package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		p := []*TreeNode{}
		for i := 0; i < len(q); i++ {
			res = append(res, q[i].Val)
			if q[i].Left != nil {
				p = append(p, q[i].Left)
			}
			if q[i].Right != nil {
				p = append(p, q[i].Right)
			}
		}
		q = p
	}
	return res
}
