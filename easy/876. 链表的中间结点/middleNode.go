package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head.Next
	}
	slow := head
	fast := head.Next

	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	return slow
}
