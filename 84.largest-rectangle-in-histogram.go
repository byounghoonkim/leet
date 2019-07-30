/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	ret := 0

	for i:= 0 ; i < len(heights); i++ {
		min_h := heights[i]
		for j := i ; j < len(heights); j++ {
			if min_h > heights[j] {
				min_h = heights[j]
			}

			rect := (j-i+1)*min_h
			if rect > ret {
				ret = rect
			}
		}
	}

	return ret
}

