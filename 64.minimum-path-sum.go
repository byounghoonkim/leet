/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (47.52%)
 * Likes:    1457
 * Dislikes: 42
 * Total Accepted:    241.7K
 * Total Submissions: 508.1K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right which minimizes the sum of all numbers along its path.
 * 
 * Note: You can only move either down or right at any point in time.
 * 
 * Example:
 * 
 * 
 * Input:
 * [
 * [1,3,1],
 * ⁠ [1,5,1],
 * ⁠ [4,2,1]
 * ]
 * Output: 7
 * Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 * 
 * 
 */

func pathSum(min *int, grid [][]int, i, j int, temp int, cache [][]int) int { 
	m:= len(grid)
	n:= len(grid[0])

	if i == m-1 && j == n-1 {
		if *min == 0 {
			*min = temp + grid[i][j]
		}
		if t := temp + grid[i][j]; *min > t {
			*min = t
		}
		return grid[i][j]
	}

	if i >=m || j >= n {
		return -1
	}

	if c := cache[i][j]; c > -1 {
		if *min == 0 {
			*min = temp + c 
		}
		if t := temp + c; *min > t {
			*min = t
		}
		return  c
	}

	a := pathSum(min, grid, i+1, j, temp+grid[i][j], cache)
	b := pathSum(min, grid, i, j+1, temp+grid[i][j], cache)

	r := 0
	if a == -1 {
		r = b
	}  else {
		if b == -1 {
			r = a
		} else {
			if a > b {
				r = b
			} else {
				r = a
			}
		}
	}


	cache[i][j] = r + grid[i][j]
	return r + grid[i][j]
}

func minPathSum(grid [][]int) int {
	ret := 0

	m:= len(grid)
	n:= len(grid[0])

	cache := make([][]int, m)
	for i,_:= range(cache) {
		cache[i] = make([]int, n)
		for j, _:= range(cache[i]) {
			cache[i][j] = -1
		}
	}

	pathSum(&ret, grid, 0,0,0, cache)

	return ret
    
}

