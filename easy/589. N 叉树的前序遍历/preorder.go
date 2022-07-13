package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var dfs func(node *Node)
	ans := []int{}
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, n := range node.Children {
			dfs(n)
		}
	}
	dfs(root)
	return ans
}
