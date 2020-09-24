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

func middleNode(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	middle := length / 2 + 1
	count := 0
	for head != nil {
		count++
		if count == middle {
			return head
		}
		head = head.Next
	}

	return head
}