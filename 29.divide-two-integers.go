/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */
func divide(dividend int, divisor int) int {
	// check overflow case ans return max int32
	if (-1<<31 == dividend) && (-1 == divisor) {
		return  int((1<<31)-1)
	}

	return int(int64(dividend) / int64(divisor))
}

