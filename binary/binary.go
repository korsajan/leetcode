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

func missingNumber(nums []int) int {
	var found = len(nums)
	for i := 0; i < len(nums); i++ {
		found ^= i
		found ^= nums[i]
	}
	return found
}

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a ^= b
		b = carry << 1
	}
	return a
}

func reverseBits(num uint32) uint32 {
	pos := 32 - 1
	var reverse = uint32(0)
	for pos >= 0 && num != 0 {
		if num&1 == 1 {
			reverse = reverse | (1 << pos)
		}
		num >>= 1
		pos--
	}

	return reverse
}
