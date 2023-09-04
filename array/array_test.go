package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tableTest = []struct {
		nums   []int
		target int
		out    []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, out: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, out: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, out: []int{0, 1}},
	}

	for i, test := range tableTest {
		res := twoSum(test.nums, test.target)
		if !equalArray(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func equalArray(a1, a2 sort.IntSlice) bool {
	a1.Sort()
	a2.Sort()
	return reflect.DeepEqual(a1, a2)
}
