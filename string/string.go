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

func isAnagram(s string, t string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		_, load := m[t[i]]
		if !load {
			return false
		}
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
