package main


/**
sliding window
time complexity: O(n)
space complexity: O(128)
start = max(start, last[s[i]]+1)
ans = max(ans, i - start + 1)
 */
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	idx := [128]int{}
	for k, _ := range idx {
		idx[k] = -1
	}

	for i, j := 0,0; j < n; j++ {
		i = max(i, idx[s[j]] + 1)
		ans = max(ans, j - i + 1)
		idx[s[j]] = j
	}
	return ans
}

func max (x , y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := " "
	lengthOfLongestSubstring(s)
}