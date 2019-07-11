/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 *
 * https://leetcode.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (31.42%)
 * Likes:    892
 * Dislikes: 115
 * Total Accepted:    182.8K
 * Total Submissions: 581.9K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * Given a set of non-overlapping intervals, insert a new interval into the
 * intervals (merge if necessary).
 * 
 * You may assume that the intervals were initially sorted according to their
 * start times.
 * 
 * Example 1:
 * 
 * 
 * Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * Output: [[1,5],[6,9]]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * Output: [[1,2],[3,10],[12,16]]
 * Explanation: Because the new interval [4,8] overlaps with
 * [3,5],[6,7],[8,10].
 * 
 * NOTE: input types have been changed on April 15, 2019. Please reset to
 * default code definition to get new method signature.
 * 
 */

func tryMerge(a , b []int) (bool, []int) {

	if a[0] > b[1] || a[1] < b[0] {
		return false, nil
	}

	i := a[0]
	j := a[1]

	if a[0] > b[0] {
		i = b[0]
	}

	if a[1] < b[1] {
		j = b[1]
	}

	return true, []int{i,j}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := [][]int{}
	l := len(intervals)

	if l == 0 {
		return append(ret, newInterval)
	}

	usedNewInterval := false
	for _, i := range(intervals) {
		if ok, m1 := tryMerge(i, newInterval); ok {
			usedNewInterval = true
			if len(ret) > 0 {
				if ok, m2 := tryMerge(ret[len(ret)-1], m1); ok {
					ret[len(ret) -1] = m2
				} else {
					ret = append(ret, m1)
				}
			} else {
				ret = append(ret, m1)
			}
		} else {
			if false == usedNewInterval && i[0] > newInterval[1] {
				usedNewInterval = true
				ret = append(ret, newInterval) 
			}
			ret = append(ret, i) 
		}
	}

	if false == usedNewInterval {
		ret = append(ret, newInterval) 
	}

	return ret

}

