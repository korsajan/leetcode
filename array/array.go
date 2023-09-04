package array

func twoSum(nums []int, target int) []int {
	var kv = make(map[int]int)
	for i, n := range nums {
		kv[n] = i
	}

	for i, n := range nums {
		j, ok := kv[target-n]
		if ok && i != j {
			return []int{i, j}
		}
	}

	return []int{}
}
