package binary

import (
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	var tableTest = []struct {
		num uint32
		out int
	}{
		{num: 11, out: 3},
		{num: 128, out: 1},
	}

	for i, test := range tableTest {
		res := hammingWeight(test.num)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestCountBits(t *testing.T) {
	var tableTest = []struct {
		num int
		out []int
	}{
		{num: 2, out: []int{0, 1, 1}},
		{num: 5, out: []int{0, 1, 1, 2, 1, 2}},
	}

	for i, test := range tableTest {
		res := countBits(test.num)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestMissingNumber(t *testing.T) {
	var tableTest = []struct {
		nums []int
		out  int
	}{
		{nums: []int{3, 0, 1}, out: 2},
		{nums: []int{0, 1}, out: 2},
		{nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, out: 8},
	}

	for i, test := range tableTest {
		res := missingNumber(test.nums)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}

func TestGetSum(t *testing.T) {
	var tableTest = []struct {
		a, b int
		out  int
	}{
		{
			a:   1,
			b:   2,
			out: 3,
		},
		{
			a:   2,
			b:   3,
			out: 5,
		},
	}

	for i, test := range tableTest {
		res := getSum(test.a, test.b)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}
