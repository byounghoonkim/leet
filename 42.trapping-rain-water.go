/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 *
 * https://leetcode.com/problems/trapping-rain-water/description/
 *
 * algorithms
 * Hard (43.53%)
 * Likes:    3742
 * Dislikes: 68
 * Total Accepted:    303.6K
 * Total Submissions: 697.4K
 * Testcase Example:  '[0,1,0,2,1,0,1,3,2,1,2,1]'
 *
 * Given n non-negative integers representing an elevation map where the width
 * of each bar is 1, compute how much water it is able to trap after raining.
 *
 *
 * The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
 * In this case, 6 units of rain water (blue section) are being trapped. Thanks
 * Marcos for contributing this image!
 *
 * Example:
 *
 *
 * Input: [0,1,0,2,1,0,1,3,2,1,2,1]
 * Output: 6
 *
 */
func trap(height []int) int {

	ret := 0

	for i, h := range height {

		l := 0
		r := 0
		for _, w := range height[:i] {
			if w > h && w > l {
				l = w
			}
		}
		for _, w := range height[i+1:] {
			if w > h && w > r {
				r = w
			}
		}

		if l > h && r > h {
			if l > r {
				ret += r - h
			} else {
				ret += l - h
			}
		}
	}

	return ret
}

