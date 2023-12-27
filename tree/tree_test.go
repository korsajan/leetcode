package tree

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	var testCases = []struct {
		in  *TreeNode
		out int
	}{
		{
			in: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Right: &TreeNode{Val: 15},
					Left:  &TreeNode{Val: 7},
				},
			},
			out: 3,
		},
		{
			in:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			out: 2,
		},
	}

	for i, test := range testCases {
		if depth := maxDepth(test.in); depth != test.out {
			t.Errorf("test %d failed got: %d want: %d", i+1, test.out, depth)
		}
	}
}

func TestIsSameTree(t *testing.T) {
	var testCases = []struct {
		p, q *TreeNode
		out  bool
	}{
		{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			out: true,
		},
		{
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			out: false,
		},
		{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			out: false,
		},
	}

	for i, test := range testCases {
		if same := isSameTree(test.p, test.q); same != test.out {
			t.Errorf("test %d failed got: %t want: %t", i+1, test.out, same)
		}
	}
}
