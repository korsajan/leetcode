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

func maxProfit(prices []int) int {
	var dayProfit = make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dayProfit[i] = make([]int, 2)
	}

	dayProfit[0][0] = -prices[0] // diff price
	dayProfit[0][1] = 0          // profit

	for i := 1; i < len(prices); i++ {
		dayProfit[i][0] = max(dayProfit[i-1][0], -prices[i])
		dayProfit[i][1] = max(dayProfit[i-1][1], dayProfit[i-1][0]+prices[i])
	}
	profit := max(dayProfit[len(prices)-1][0], dayProfit[len(prices)-1][1])
	return profit
}

func containsDuplicate(nums []int) bool {
	var filter = make(map[int]struct{})
	for _, n := range nums {
		if _, ok := filter[n]; ok {
			return ok
		}
		filter[n] = struct{}{}
	}
	return false
}

func productExceptSelf(nums []int) []int {
	var n = 1
	var sum = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum[i] = n
		n *= nums[i]
	}

	n = 1
	for i := len(nums) - 1; i >= 0; i-- {
		sum[i] *= n
		n *= nums[i]
	}

	return sum
}

const (
	intSize = 32 << (^uint(0) >> 63)
	maxInt  = 1<<(intSize-1) - 1
	minInt  = -1 << (intSize - 1)
)

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	var bestSum, currSum = minInt, 0
	for _, n := range nums {
		currSum = max(currSum+n, n)
		if currSum > bestSum {
			bestSum = currSum
		}
	}

	return bestSum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
