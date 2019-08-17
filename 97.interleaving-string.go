/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 *
 * https://leetcode.com/problems/interleaving-string/description/
 *
 * algorithms
 * Hard (28.62%)
 * Likes:    860
 * Dislikes: 48
 * Total Accepted:    118.5K
 * Total Submissions: 414.1K
 * Testcase Example:  '"aabcc"\n"dbbca"\n"aadbbcbcac"'
 *
 * Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and
 * s2.
 *
 * Example 1:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
 * Output: false
 *
 *
 */

func check(s1, s2, s3 string) bool {
	ls1 := len(s1)
	ls2 := len(s2)
	ls3 := len(s3)

	if !(ls3 == ls1+ls2) {
		return false
	}

	if ls3 == 0 {
		if ls1 > 0 || ls2 > 0 {
			return false
		} else {
			return true
		}
	}

	if ls1 == 0 && ls2 == 0 {
		return false
	}

	if ls1 > 0 && s1[0] == s3[0] {
		if check(s1[1:], s2, s3[1:]) {
			return true
		}
	}

	if ls2 > 0 && s2[0] == s3[0] {
		if check(s1, s2[1:], s3[1:]) {
			return true
		}
	}

	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {

	return check(s1, s2, s3)

}

