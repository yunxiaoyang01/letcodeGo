package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := zigzagLevelOrder(root)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%d\n",res[i])
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root==nil{
		return [][]int{}
	}
	var res [][]int
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	isLeft := true
	for len(queue) > 0 {
		l := len(queue)
		ans := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if isLeft {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}
		}
		isLeft = !isLeft
		res = append(res, ans)
		queue = queue[l:]
	}
	return res
}
