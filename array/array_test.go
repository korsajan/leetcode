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
		if !equalArrays(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestContainsDuplicate(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  bool
	}{
		{nums: []int{1, 2, 3, 1}, out: true},
		{nums: []int{1, 2, 3, 4}, out: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, out: true},
	}

	for i, test := range tableTest {
		res := containsDuplicate(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{7, 1, 5, 3, 6, 4}, out: 5},
		{nums: []int{7, 6, 4, 3, 1}, out: 0},
	}

	for i, test := range tableTest {
		res := maxProfit(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  []int
	}{
		{nums: []int{1, 2, 3, 4}, out: []int{24, 12, 8, 6}},
		{nums: []int{-1, 1, 0, -3, 3}, out: []int{0, 0, 9, 0, 0}},
	}

	for i, test := range tableTest {
		res := productExceptSelf(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestMaxSubArray(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, out: 6},
		{nums: []int{1}, out: 1},
		{nums: []int{5, 4, -1, 7, 8}, out: 23},
		{nums: []int{-2, -1}, out: -1},
		{nums: []int{-3, -2, -2, -3}, out: -2},
	}

	for i, test := range tableTest {
		res := maxSubArray(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestMaxProduct(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{2, 3, -2, 4}, out: 6},
		{nums: []int{-2, 0, -1}, out: 0},
		{nums: []int{-2}, out: -2},
		{nums: []int{-3, -1, -1}, out: 3},
		{nums: []int{-2, 3, -4}, out: 24},
		{nums: []int{0, 2}, out: 2},
		{nums: []int{-4, -3, -2}, out: 12},
	}

	for i, test := range tableTest[:1] {
		res := maxProduct(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestFindMin(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{3, 4, 5, 1, 2}, out: 1},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, out: 0},
		{nums: []int{11, 13, 15, 17}, out: 11},
		{nums: []int{3, 0, 1, 2}, out: 0},
	}

	for i, test := range tableTest {
		res := findMin(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestSearch(t *testing.T) {
	var tableTest = []struct {
		nums   []int
		target int
		out    int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, out: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, out: -1},
		{nums: []int{1, 2, 3}, target: 2, out: 1},
		{nums: []int{1}, target: 0, out: -1},
	}

	for i, test := range tableTest {
		res := search(test.nums, test.target)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestMaxArea(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, out: 49},
		{nums: []int{1, 1}, out: 1},
	}

	for i, test := range tableTest {
		res := maxArea(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestThreeSum(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  [][]int
	}{
		{nums: []int{-1, 0, 1, 2, -1, -4}, out: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{0, 1, 1}, out: [][]int{}},
		{nums: []int{0, 0, 0}, out: [][]int{{0, 0, 0}}},
	}

	for i, test := range tableTest {
		res := threeSum(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func equalArrays(a1, a2 sort.IntSlice) bool {
	a1.Sort()
	a2.Sort()
	return reflect.DeepEqual(a1, a2)
}
