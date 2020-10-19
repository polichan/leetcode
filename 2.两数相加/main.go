package main

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{6, &ListNode{4, nil}}}
	addTwoNumbers(l1, l2)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{0, nil}
	res := list
	tmp := 0
	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		list = list.Next
	}
	return res.Next
}