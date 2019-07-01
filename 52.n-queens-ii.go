/*
 * @lc app=leetcode id=52 lang=golang
 *
 * [52] N-Queens II
 *
 * https://leetcode.com/problems/n-queens-ii/description/
 *
 * algorithms
 * Hard (52.37%)
 * Likes:    285
 * Dislikes: 116
 * Total Accepted:    101.8K
 * Total Submissions: 194.3K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return the number of distinct solutions to the n-queens
 * puzzle.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: 2
 * Explanation: There are two distinct solutions to the 4-queens puzzle as
 * shown below.
 * [
 * [".Q..",  // Solution 1
 * "...Q",
 * "Q...",
 * "..Q."],
 *
 * ["..Q.",  // Solution 2
 * "Q...",
 * "...Q",
 * ".Q.."]
 * ]
 *
 *
 */

func canPutQueen(current [][]byte, i, j int) bool {
	// 주어진 i, j 에 Queen을 높을 수 있는지 반환하는 메소드
	n := len(current)
	for a := 0; a < n; a++ {
		// 같은 row에 퀸이 있는지 확인
		if current[a][j] == 'Q' {
			return false
		}

		// 같은 컬럼에 퀀이 있는지 확인
		if current[i][a] == 'Q' {
			return false
		}

		// 우하단 사선에 여왕이 있는지 확인
		if i+a < n && j+a < n {
			if current[i+a][j+a] == 'Q' {
				return false
			}
		}

		// 좌상
		if i-a >= 0 && j-a >= 0 {
			if current[i-a][j-a] == 'Q' {
				return false
			}
		}

		// 좌하
		if i-a >= 0 && j+a < n {
			if current[i-a][j+a] == 'Q' {
				return false
			}
		}

		// 우상
		if i+a < n && j-a >= 0 {
			if current[i+a][j-a] == 'Q' {
				return false
			}
		}
	}

	return true
}

func helper(ret *int, current [][]byte, current_count int, current_index int) {
	n := len(current)
	max_index := n * n

	// 여왕이 n개 놓아 져서 진입하면 해법 갯수를 증가 시키고 리턴한다.
	if current_count == n {
		*ret += 1
		return
	}

	// 보드 끝까지 진행 했으면 리턴한다.
	if current_index == max_index {
		return
	}

	i := current_index / n
	j := current_index % n

	// 여왕을 놓을 수 있으면 놓고 재귀 호출 후 여왕을 제거한다.
	if true == canPutQueen(current, i, j) {
		current[i][j] = 'Q'
		helper(ret, current, current_count+1, current_index+1)
		current[i][j] = '.'
	}

	// 여왕을 놓지 않은 상태에서 앞으로 진행한다.
	helper(ret, current, current_count, current_index+1)
}

func totalNQueens(n int) int {
	ret := 0

	// 임시 솔루션 말판을 만든다.
	// []string 이 아니라 [][]byte 로 만드는 이유는
	// golang 에서 string 은 변경 불가 byte 배열이라
	// 여왕을 놓았다 지웠다 하기 번거롭기 때문이다.
	current := make([][]byte, n)
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j, _ := range row {
			row[j] = '.'
		}

		current[i] = row

	}

	helper(&ret, current, 0, 0)

	return ret

	// 이 알고리즘은
	// n 에 대해 n^2 만큼 루프를 돌며 각 루프 마다 최대 4n 만큼 루프를 돌게 된다.
	// 최악의 경우 4n^3 정도의 수행을 해야 하므로 O(n^3) 으로 표현할 수 있겠다.
}

