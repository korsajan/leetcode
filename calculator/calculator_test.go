package calculator

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			in:  "1 + 1",
			out: 2,
		},
		{
			in:  " 2-1 + 2 ",
			out: 3,
		},
		{
			in:  "(1+(6+8))",
			out: 15,
		},
		{
			in:  "(1+(4+5+2)-3)+(6+8)",
			out: 23,
		},
	}
	for i, test := range tests {
		res := calculate(test.in)
		if res != test.out {
			t.Errorf("test %d excpected : %d got : %d",
				i+1, test.out, res)
		}
	}
}
