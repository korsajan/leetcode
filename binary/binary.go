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
