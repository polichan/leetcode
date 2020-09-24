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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	// 快慢指针，快慢指针都指向链头，快指针先走 K-1 ，然后快慢指针同时走，如果快针到了链尾，则代表慢指针已经走到了倒数第K个位置
	slow, fast := head, head
	for i := 0; i< k - 1; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}