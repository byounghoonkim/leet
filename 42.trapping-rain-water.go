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

	// 모든 요소를 순회 하면서 해당 요소에 물이 담길수 있는지 체크 하고
	// 담길 수 있다면 용량을 측정해 더한다.
	for i, h := range height {

		l := 0 // 왼편에서 제일 높은 벽 높이
		r := 0 // 오른편에서 제일 높은 벽 높이

		// i 를 중심으로 왼편에서 제일 높은 벽을 찾는다.
		for _, w := range height[:i] {
			if w > h && w > l {
				l = w
			}
		}

		// i 를 중심으로 오른편에 제일 높은 벽을 찾는다.
		for _, w := range height[i+1:] {
			if w > h && w > r {
				r = w
			}
		}

		// 양측에서 찾은 벽 중 낮은 벽 높이와 i 의 높이 h 의 차이를 그 요소의
		// 용량으로 더한다.
		if l > h && r > h {
			if l > r {
				ret += r - h
			} else {
				ret += l - h
			}
		}
	}

	return ret

	// 이 알고리즘은 height 의 개수 n 의 모든 요소를 순회하고
	// 해당 순회 시 양측으로 한번씩 순회를 하므로
	// n*n을 순회 한다. O(n^2)
}

