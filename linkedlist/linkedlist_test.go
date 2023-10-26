package linkedlist

import (
	"reflect"
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
