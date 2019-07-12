/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 *
 * https://leetcode.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (47.20%)
 * Likes:    480
 * Dislikes: 81
 * Total Accepted:    140.6K
 * Total Submissions: 297.8K
 * Testcase Example:  '3'
 *
 * Given a positive integer n, generate a square matrix filled with elements
 * from 1 to n^2 in spiral order.
 * 
 * Example:
 * 
 * 
 * Input: 3
 * Output:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 * 
 * 
 */
func generateMatrix(n int) [][]int {

	ret := make([][]int, n)

	// n*n 슬라이스를 생성한다.
	for i, _ := range(ret) {
		ret[i] = make([]int, n)
	}

	count := 1
	// 최 외곽부터 한열씩 안으로 진입하기 위한 루프를 돈다.
	for x:= 0 ; x <= n/2 ; x++ {

		// 상단 행 채움
		for j:= x; j < n-1-x ; j++ {
			ret[x][j] = count
			count++
		}

		// 우측 열
		for i:=x; i < n-1-x ; i++ {
			ret[i][n-1-x] = count
			count++
		}

		// 하단 행
		for j:= n-1-x; j > x ; j-- {
			ret[n-1-x][j] = count
			count++
		}

		// 좌측 열
		for i:=n-1-x; i > x ; i-- {
			ret[i][x] = count
			count++
		}
	}

	//n 이 홀수 이면 마지막 한 칸은 채워지지 않으므로 여기서 채워준다.
	if 1 == n%2 {
		ret[n/2][n/2] = count
	}

	return ret

	// 이 알고리즘은 n 의 개수 n^2 만큼의 아이템에 루프를 돌아야 하므로
	// O(n^2)의 수행시간 복잡도를 가진다.
}

