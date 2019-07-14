/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (48.14%)
 * Likes:    1683
 * Dislikes: 118
 * Total Accepted:    303.7K
 * Total Submissions: 630.5K
 * Testcase Example:  '3\n2'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 * 
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 * 
 * How many possible unique paths are there?
 * 
 * 
 * Above is a 7 x 3 grid. How many possible unique paths are there?
 * 
 * Note: m and n will be at most 100.
 * 
 * Example 1:
 * 
 * 
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation:
 * From the top-left corner, there are a total of 3 ways to reach the
 * bottom-right corner:
 * 1. Right -> Right -> Down
 * 2. Right -> Down -> Right
 * 3. Down -> Right -> Right
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: m = 7, n = 3
 * Output: 28
 * 
 */

func helper(i, j int, cache [][]int) int {
	if i == 1 || j == 1 {
		return 1 
	}

	if cache[i][j] > 0 {
		return cache[i][j] 
	}

	temp := 0
	temp += helper(i, j-1, cache)
	temp += helper(i-1,j, cache)

	cache[i][j] = temp

	return temp 

}

func uniquePaths(m int, n int) int {
	cache := make([][]int, m+1)
	for i, _ := range(cache) {
		cache[i] = make([]int, n+1)
	}

	return helper(m, n, cache)

    
}

