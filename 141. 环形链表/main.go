package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool, 0)
	for head != nil {
		if !m[head] {
			m[head] = true
		}else {
			return true
		}
		head = head.Next
	}
	return false
}