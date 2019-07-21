/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (38.45%)
 * Likes:    2252
 * Dislikes: 34
 * Total Accepted:    183.2K
 * Total Submissions: 476.3K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two words word1 and word2, find the minimum number of operations
 * required to convert word1 to word2.
 * 
 * You have the following 3 operations permitted on a word:
 * 
 * 
 * Insert a character
 * Delete a character
 * Replace a character
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation: 
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation: 
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 * 
 * 
 */
func minDistance(word1 string, word2 string) int {

	l1 := len(word1)+1
	l2 := len(word2)+1

	m := make([][]int, l1)
	for i, _:=range(m) {
		m[i] = make([]int, l2)
	}
	
	for i := 0 ; i < l1 ; i ++ {
		m[i][0] = i
	}

	for j := 0 ; j < l2 ; j ++ {
		m[0][j] = j
	}

	for i := 1; i < l1 ; i ++ {
		for j := 1; j < l2; j++ {
			if word1[i-1] == word2[j-1] {
				m[i][j] = m[i-1][j-1]
			} else {
				min := m[i-1][j-1]
				if min > m[i][j-1]{
					min = m[i][j-1]
				}
				if min > m[i-1][j]{
					min = m[i-1][j]
				}
				m[i][j] = min + 1
			}
		}
	}

	return m[l1-1][l2-1]
}

