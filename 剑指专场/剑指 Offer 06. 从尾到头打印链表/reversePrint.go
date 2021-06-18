package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	resReverse := []int{}
	for i:=len(res)-1; i >=0 ; i-- {
		resReverse = append(resReverse,res[i])
	}
	return resReverse
}


// 递归解法 直接翻转
func reversePrintV2(head *ListNode) []int {
	res := []int{}
	reverse := func(node *ListNode){}
	reverse = func(node *ListNode){
		if node==nil {
			return 
		}
		reverse(node.Next)
		res = append(res, node.Val)
	}
	reverse(head)
	return res
}