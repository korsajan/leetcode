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

func HasCycle(head *ListNode) bool {
	seen := make(map[*ListNode]bool)

	first := head
	for n := first; n != nil; n = n.Next {
		if seen[n] {
			return true
		}
		seen[n] = true
	}
	return false
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		res  *ListNode
		prev *ListNode
		node *ListNode
	)

	var over bool
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		if over {
			sum++
		}

		if sum >= 10 {
			over = true
		} else {
			over = false
		}

		node = &ListNode{Val: sum % 10, Next: nil}
		if res == nil {
			res = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	if over {
		node.Next = &ListNode{Val: 1, Next: nil}
	}
	return res
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
