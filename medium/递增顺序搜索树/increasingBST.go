package main
type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}
/*
	题解：由题目意思可以知道这道题是中序遍历一个二叉搜索树，然后将这个遍历的结果组合成一个
	只有有节点的树
	解题思路：
	第一步：直接先来一个中序遍历先把这颗树的值都遍历出来，存在一个数组里面。
	第二步：定义一个根节点，然后将根节点的值赋给一个临时的节点。
	第三步：开始循环我们刚才遍历出来的数组，将临时节点的右节点val赋值
	最后 把根节点的右节点作为结果返回
*/
func increasingBST(root *TreeNode) *TreeNode {
	var vals []int
	var inOrder func(*TreeNode)
	inOrder  = func(tn *TreeNode) {
		if tn !=nil{
			inOrder(tn.Left)
			vals =append(vals, tn.Val)
			inOrder(tn.Right)
		}
	}
	inOrder(root)
	result := &TreeNode{}
	currentNode := result
	for _,val := range vals{
		currentNode.Right = &TreeNode{Val: val}
		currentNode = currentNode.Right
	}
	return result.Right
}