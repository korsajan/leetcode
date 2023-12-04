package string

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		in  string
		out int
	}{
		{in: "abcabcbb", out: 3},
		{in: "bbbb", out: 1},
		{in: "pwwkew", out: 3},
	}
	for i, test := range tests {
		res := lengthOfLongestSubstring(test.in)
		if !reflect.DeepEqual(test.out, res) {
			t.Errorf("test %d failed ex: %v got: %v\n", i+1, test.out, res)
		}
	}
}
