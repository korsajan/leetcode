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
