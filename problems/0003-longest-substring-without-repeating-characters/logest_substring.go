package problem0003

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
func lengthOfLongestSubstring(s string) int {
	locations := make(map[rune]int)
	maxLen, left := 0, 0

	for i, c := range s {
		l, ok := locations[c]
		if ok && l >= left {
			if maxLen < i-left {
				maxLen = i - left
			}
			left = l + 1
		}
		locations[c] = i
	}

	if maxLen < len(s)-left {
		maxLen = len(s) - left
	}
	return maxLen
}
