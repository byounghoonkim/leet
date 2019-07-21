/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 *
 * https://leetcode.com/problems/set-matrix-zeroes/description/
 *
 * algorithms
 * Medium (40.32%)
 * Likes:    1147
 * Dislikes: 204
 * Total Accepted:    220.6K
 * Total Submissions: 547.1K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given a m x n matrix, if an element is 0, set its entire row and column to
 * 0. Do it in-place.
 * 
 * Example 1:
 * 
 * 
 * Input: 
 * [
 * [1,1,1],
 * [1,0,1],
 * [1,1,1]
 * ]
 * Output: 
 * [
 * [1,0,1],
 * [0,0,0],
 * [1,0,1]
 * ]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 
 * [
 * [0,1,2,0],
 * [3,4,5,2],
 * [1,3,1,5]
 * ]
 * Output: 
 * [
 * [0,0,0,0],
 * [0,4,5,0],
 * [0,3,1,0]
 * ]
 * 
 * 
 * Follow up:
 * 
 * 
 * A straight forward solution using O(mn) space is probably a bad idea.
 * A simple improvement uses O(m + n) space, but still not the best
 * solution.
 * Could you devise a constant space solution?
 * 
 * 
 */
func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])

	if m < 1 { return }
	if n < 1 { return }

	// m*n 또는 m+n개의 배열을 준비하고 해당 행열이 0이 되어야 하는지
	// 기억해 놓으면 되지만, 문제의 팔로업 부분을 읽어 보면 
	// 더 좋은 솔루션이 있는지 요청하고 있다.
	// 이에 첫 행과 열이 0인지만 기악하는 변수를 두개 두고
	// 이를 먼저 검색한 후
	fr := false
	fc := false
	
	for j := 0 ; j < n; j++{
		if 0 == matrix[0][j] {
			fr = true
			break
		}
	}

	for i := 0 ; i < m; i++{
		if 0 == matrix[i][0] {
			fc = true
			break
		}
	}

	// 나머지 1 이상의 행열들을 검색하면서 0이면 첫번쩌 행열에 0을 표시 하는 방법을 취한다.
	for i := 1 ; i < m; i++{
		for j := 1; j < n; j++ {
			if 0 == matrix[i][j] {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 그리고 다시 돌면서 0번째 행열에 0이 표시 되어 있다면 나머지 행열을 0으로 채우고
	for i := 1 ; i < m; i++{
		for j := 1; j < n; j++ {
			if 0 ==matrix[i][0] {
				matrix[i][j] = 0
			}
			if 0 == matrix[0][j] {
				matrix[i][j] = 0
			}
		}
	}

	// 최종 첫번째 행열이 0이라면 해당 행열도 0으로 채운다.
	if fr {
		for j := 0 ; j < n; j++{
			matrix[0][j] = 0
		}
	}

	if fc {
		for i := 0 ; i < m; i++{
			matrix[i][0] = 0
		}
	}


	// 이 알고리즘은 m*n 행열 전체를 정확히 두번 스캔 하므로
	// 수행 시간 복잡도는 2m*n 으로 O(m*n) 이다.
	// 공간 복잡도는 m*n의 크기와 상관 없이 2개의 불리언 변수만 추가 되므로
	// O(1) 의 공간 복잡도를 가진다.
}

