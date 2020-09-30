package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first := dummy
	second := dummy.Next
	for  {
		if second.Val == val {
			first.Next = second.Next
			break
		}
		first = first.Next
		second = second.Next
	}
	return dummy.Next
}

