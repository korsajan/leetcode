package linkedlist

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		in  *ListNode
		out *ListNode
	}{
		{
			in:  SliceToList(1, 2, 3, 4, 5),
			out: SliceToList(5, 4, 3, 2, 1),
		},
		{
			in:  SliceToList(1, 2),
			out: SliceToList(2, 1),
		},
		{
			in:  SliceToList(),
			out: SliceToList(),
		},
	}
	for i, test := range tests {
		l := ReverseList(test.in)
		if !reflect.DeepEqual(l, test.out) {
			t.Errorf("test %d failed ex: %v want: %v", i+1, test.out, l)
		}
	}
}

func TestAddTwoNumbers(t *testing.T) {
	testTable := []struct {
		l1, l2, out *ListNode
	}{
		{
			l1:  SliceToList(2, 4, 3),
			l2:  SliceToList(5, 6, 4),
			out: SliceToList(7, 0, 8),
		},
		{
			l1:  SliceToList(0),
			l2:  SliceToList(0),
			out: SliceToList(0),
		},
		{
			l1:  SliceToList(9, 9, 9, 9, 9, 9, 9),
			l2:  SliceToList(9, 9, 9, 9),
			out: SliceToList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}
	for i, test := range testTable {
		l := addTwoNumbers(test.l1, test.l2)
		if !reflect.DeepEqual(l, test.out) {
			t.Errorf("test %d failed ex: %v want: %v", i+1,
				ToString(test.out),
				ToString(l))
		}
	}
}

func ToString(l *ListNode) string {
	var b strings.Builder
	b.WriteRune('{')
	n := l
	for n != nil {
		s := strconv.Itoa(n.Val)
		b.WriteString(s)
		if n.Next != nil {
			b.WriteString("->")
		}
		n = n.Next
	}
	b.WriteRune('}')
	return b.String()
}
