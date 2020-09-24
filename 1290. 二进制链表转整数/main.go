package main

func main() {
	
}

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


func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = result << 1 | head.Val
		head = head.Next
	}
	return result
}

