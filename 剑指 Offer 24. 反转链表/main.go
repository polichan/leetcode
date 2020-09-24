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

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	pre := head
	for pre != nil {
		temp := pre.Next
		pre.Next = cur
		cur = pre
		pre = temp
	}
	return cur
}