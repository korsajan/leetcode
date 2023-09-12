package binary

func hammingWeight(num uint32) int {
	var w int
	for num > 0 {
		if num&1 == 1 {
			w++
		}
		num = num >> 1
	}
	return w
}

func countBits(n int) []int {
	var out = make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		out = append(out, hammingWeight(uint32(i)))
	}
	return out
}
