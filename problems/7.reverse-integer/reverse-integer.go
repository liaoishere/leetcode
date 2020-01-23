package problem7

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	answer := 0
	for x > 0 {
		answer = x%10 + answer*10
		x /= 10
	}
	if answer > math.MaxInt32 || answer < -math.MaxInt32-1 {
		return 0
	}
	return answer * sign
}
