package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	flag := true
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			res[i] = append(res[i], q[j].Val)
			if q[j].Left != nil {
				p = append(p, q[j].Left)
			}
			if q[j].Right != nil {
				p = append(p, q[j].Right)
			}
		}
		q = p
		if !flag {
			flag = true
			res[i] = reverseInt(res[i])
		} else {
			flag = false
		}
	}
	return res
}

func reverseInt(val []int) []int {
	res := []int{}
	for i := len(val) - 1; i >= 0; i-- {
		res = append(res, val[i])
	}
	return res
}
