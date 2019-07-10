/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 *
 * https://leetcode.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (35.86%)
 * Likes:    2211
 * Dislikes: 168
 * Total Accepted:    366.2K
 * Total Submissions: 1M
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * Given a collection of intervals, merge all overlapping intervals.
 * 
 * Example 1:
 * 
 * 
 * Input: [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into
 * [1,6].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [[1,4],[4,5]]
 * Output: [[1,5]]
 * Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */

func sort(intervals [][]int) {
	l := len(intervals)
	for i:= 0 ; i < l ; i++ {
		for j:= 0 ; j < l - i -1; j++{
			if intervals[j][0] > intervals[j+1][0] {
			intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
}

func merge(intervals [][]int) [][]int {

	ret := [][]int{}

	sort(intervals)

	for _, i := range intervals {
		merged := false
		for _, r := range ret {
			if i[0] <= r[1] {
				if i[1] > r[1] {
					r[1] = i[1]
				}
				merged = true
			}
		}

		if merged == false {
			ret = append(ret, i)
		}
	}

	return ret
    
}

