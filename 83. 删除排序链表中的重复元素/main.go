package main

func main() {
	
}
type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prevNode := head
	currNode := head.Next

	for currNode != nil {
		if prevNode.Val == currNode.Val {
			prevNode.Next = currNode.Next
		}else {
			prevNode = currNode
		}
		currNode = currNode.Next
	}
	return head
}