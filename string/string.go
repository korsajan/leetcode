package string

func lengthOfLongestSubstring(s string) int {
	visited := make([]bool, 256)
	low, high, seq := 0, 0, 0
	for high < len(s) {
		if visited[s[high]] {
			for visited[s[high]] {
				visited[s[low]] = false
				low++
			}
		}
		visited[s[high]] = true
		seq = mmax(seq, high-low+1)
		high++
	}
	return seq
}

func mmax(i, j int) int {
	return max(i, j)
}
