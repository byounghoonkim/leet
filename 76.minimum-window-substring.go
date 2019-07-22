/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 *
 * https://leetcode.com/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (31.33%)
 * Likes:    2467
 * Dislikes: 165
 * Total Accepted:    252.5K
 * Total Submissions: 806K
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * Given a string S and a string T, find the minimum window in S which will
 * contain all the characters in T in complexity O(n).
 * 
 * Example:
 * 
 * 
 * Input: S = "ADOBECODEBANC", T = "ABC"
 * Output: "BANC"
 * 
 * 
 * Note:
 * 
 * 
 * If there is no such window in S that covers all characters in T, return the
 * empty string "".
 * If there is such window, you are guaranteed that there will always be only
 * one unique minimum window in S.
 * 
 * 

 "ADOBECODEBANC", T = "ABC"
  1         1
	 1     1   
	   1      1

 */

func minMax(t []int) (int, int) {
	min, max := t[0], t[0]

	for _, a := range(t) {
		if min > a {
			min = a
		}
		if max < a {
			max = a
		}
	}

	return min, max
}

func spanLen(t []int) int {
	min, max := minMax(t)
	return max - min
}

func findWindow(m [][]int, t []int, min *int, r *[]int ) {
	if 0 == len(m) {
		if *min == -1 || *min > spanLen(t) {
			*min = spanLen(t)
			*r = make([]int, len(t))
			copy(*r,t)
		}
		return
	}

	for _, a := range(m[0]) {
		findWindow(m[1:], append(t, a), min, r)
	}
}

func minWindow(s string, t string) string {

	m := make([][]int, len(t))


	for i, sc := range(s) {
		for j, tc := range(t) {
			if sc == tc {
				m[j] = append(m[j], i)
			}
		}
	}


	for _, a := range(m) {
		if len(a) == 0 {
			return ""
		}
	}


	min := -1
	ret := []int{}

	findWindow(m, []int{}, &min, &ret)

	min, max := minMax(ret)

	if max+1 - min < len(m) {
		return ""
	}
	
	return s[min:max+1]
}

