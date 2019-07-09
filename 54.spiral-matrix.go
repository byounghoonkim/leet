/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (30.67%)
 * Likes:    1147
 * Dislikes: 420
 * Total Accepted:    242.1K
 * Total Submissions: 788.3K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix of m x n elements (m rows, n columns), return all elements of
 * the matrix in spiral order.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * Output: [1,2,3,6,9,8,7,4,5]
 *
 *
 * Example 2:
 *
 * Input:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 */

func spiralOrder(matrix [][]int) []int {
	ret := []int{}

	m:= len(matrix)
	if m == 0 {
		// 빈 슬라이스가 들어 오면 빈 슬라이스를 리턴한다.
		return ret
	}
	n:= len(matrix[0])

	// 초기 요소 0,0 에서 시작한다.
	i, j := 0, 0

	for ;; {
		// 최후의 하나 요소가 남으면 그 요소를 추가 하고 끝낸다.
		if m == 1 && n == 1 {
			ret=append(ret, matrix[i][j])
			break;
		}

		// 상단 열을 추가한다.
		for step:= 0 ; step < n-1 ; step++ {
			ret=append(ret, matrix[i][j])
			j++
		}

		// 우측 열을 진행 하기 전에 m 이 1이면 
		// 현재 가리키고 있는 요소만 추가 하고 끝낸다.
		if m == 1 {
			ret=append(ret, matrix[i][j])
			break;
		}

		// 우측 열을 추가한다.
		for step:= 0 ; step < m-1 ; step++ {
			ret=append(ret, matrix[i][j])
			i++	
		}

		// 하단 열을 진행 하기 전에
		// n 이 1 이면 현재 가키리는 요소만 추가하고 끝낸다.
		if n == 1 {
			ret=append(ret, matrix[i][j])
			break;
		}

		// 하단 열을 추가한다.
		for step:= 0 ; step < n-1 ; step++ {
			ret=append(ret, matrix[i][j])
			j--
		}

		// 좌측 열을 추가한다.
		for step:= 0 ; step < m-1 ; step++ {
			ret=append(ret, matrix[i][j])
			i--
		}

		// 최외측 열을 제외 하기 위해 2씩 뺀다.
		m = m - 2
		n = n - 2

		// 남은 배열이 없으면 끝낸다.
		if m <=0 || n <= 0 {
			break
		}

		// 요소 좌표에 1씩 더해 안쪽 열로 위치를 바꾼다.
		i ++
		j ++
	}

	return ret

	// 이 알고리즘은 제일 바깥쪽 열을 순회 하면서 결과를 추가하고
	// i, j 에 1씩 추가 해 안쪽 열로 진입해 4면의 열을 추가 하는 방식으로
	// 결과를 구한다.
	// 전체 순회 하는 갯수는 n 에 비례하므로 
	// 수행 시간 복잡도는 O(n) 이다.
}

