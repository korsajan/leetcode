package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var curr, prev, next *ListNode = head, nil, nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func SliceToList(arr ...int) *ListNode {
	var node *ListNode
	for i := 0; i < len(arr); i++ {
		n := &ListNode{Val: arr[i], Next: nil}
		if node == nil {
			node = n
			continue
		} else {
			h := node
			for h.Next != nil {
				h = h.Next
			}
			h.Next = n
		}
	}
	return node
}
