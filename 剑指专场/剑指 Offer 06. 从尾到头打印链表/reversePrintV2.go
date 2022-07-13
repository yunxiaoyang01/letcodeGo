package main

func reversePrintV3(head *ListNode) []int {
	reverse := func(node *ListNode) {}
	res := []int{}
	reverse = func(node *ListNode) {
		if node == nil {
			return
		}
		reverse(node.Next)
		res = append(res, node.Val)
	}
	reverse(head)
	return res
}
