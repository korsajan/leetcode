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
