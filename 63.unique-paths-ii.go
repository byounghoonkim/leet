/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.56%)
 * Likes:    912
 * Dislikes: 135
 * Total Accepted:    209.4K
 * Total Submissions: 623.9K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 * 
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 * 
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 * 
 * 
 * 
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 * 
 * Note: m and n will be at most 100.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 * 
 * 
 */


func helper(o [][]int, i, j, m, n int, cache [][]int) int {

	// 맵 외부로 탐색하는 것은 경로 없는 0 으로 리턴한다.
	if i >= m || j >= n {
		return 0
	}

	// 장애물이 있다면 0 으로 리턴한다.
	if o[i][j] == 1 {
		return 0
	}

	// 목적지에 도달했다면 1로 리턴한다.
	if i == m-1 && j == n-1 {
		return 1
	}

	// 캐시가 0 이상이라면 한번이라도 방문 했기 때문에 캐시를 활용한다.
	if c := cache[i][j]; c >= 0  {
		return c 
	}

	// 다운 , 라이트 순으로 하위 맵을 탐색한다.
	temp := 0
	temp += helper(o, i+1, j, m, n, cache)
	temp += helper(o, i, j+1, m, n, cache)

	// 캐시를 셋팅한다.
	cache[i][j] = temp

	return temp
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	// 주어진 맵의 크기를 구한다.
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 캐시를 만들고 -1로 채운다.
	// 0으로 초기화 할 경우 경로가 없는 경우와 구별되지 않아
	// 캐시의 히트 성능이 떨어진다.
	cache := make([][]int, m)
	for i, _ := range(cache) {
		cache[i] = make([]int, n)

		for j, _ := range(cache[i]) {
			cache[i][j] = -1
		}
	}

	return helper(obstacleGrid,0,0,m, n, cache)
}

