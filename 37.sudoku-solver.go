/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 *
 * https://leetcode.com/problems/sudoku-solver/description/
 *
 * algorithms
 * Hard (37.36%)
 * Likes:    903
 * Dislikes: 60
 * Total Accepted:    132.3K
 * Total Submissions: 354.1K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Write a program to solve a Sudoku puzzle by filling the empty cells.
 *
 * A sudoku solution must satisfy all of the following rules:
 *
 *
 * Each of the digits 1-9 must occur exactly once in each row.
 * Each of the digits 1-9 must occur exactly once in each column.
 * Each of the the digits 1-9 must occur exactly once in each of the 9 3x3
 * sub-boxes of the grid.
 *
 *
 * Empty cells are indicated by the character '.'.
 *
 *
 * A sudoku puzzle...
 *
 *
 * ...and its solution numbers marked in red.
 *
 * Note:
 *
 *
 * The given board contain only digits 1-9 and the character '.'.
 * You may assume that the given Sudoku puzzle will have a single unique
 * solution.
 * The given board size is always 9x9.
 *
 *
 */

func b2i(b byte) int {
	return int(b - '0')
}

func i2b(i int) byte {
	return byte(i + '0')
}

func candidates(board [][]byte, r int, c int) []int {
	ret := []int{}

	check := [9]bool{}

	for i := 0; i < 9; i++ {
		if j := board[r][i]; j != '.' {
			check[b2i(j)-1] = true
		}
		if j := board[i][c]; j != '.' {
			check[b2i(j)-1] = true
		}
	}

	for i := (r / 3) * 3; i < (r/3)*3+3; i++ {
		for j := (c / 3) * 3; j < (c/3)*3+3; j++ {
			if b := board[i][j]; b != '.' {
				check[b2i(b)-1] = true
			}
		}
	}

	for i, b := range check {
		if b == false {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func solve(board [][]byte, level int) bool {
	// level 변수는 보드의 요소 81개를 일렬로 늘러뜨렸을 때 앞으로 진행한
	// 정도를 나타낸다.
	// 0부터 시작해 80이 되면 모든 요소의 값을 찾았으므로 true 로 리턴한다.
	if level == 81 {
		return true
	}

	// level에서 r, c 인덱스를 계산해 낸다.
	r := level / 9
	c := level % 9

	// 값이 이미 있으면 항목이면 다음 값으로 넘어 간다.
	if '.' != board[r][c] {
		return solve(board, level+1)
	}

	// 해당 셀에 입력 가능한 값 목록을 얻는다.
	candi := candidates(board, r, c)

	// 각 값들을 넣고 재귀적으로 solve 함수를 호출한다.
	for _, ca := range candi {
		board[r][c] = i2b(ca)

		// 함수가 true를 리턴하면 이 함수도 true를 리턴한다.
		if true == solve(board, level+1) {
			return true
		}

		// 함수가 실패 하면 다시 . 으로 되돌려 놓는다. 이렇게 하지 않으면 마지막에
		// 찌꺼기 값이 남는다.
		board[r][c] = '.'
	}
	return false

	// 이 알고리즘은 81개의 항목에서 후보 값을 얻기 위해 30개의 항목을 검사 하고
	// 각 항목의 후보 값들의 곱 만큼 유효성 검사를 위해 30개의 항목을 검사한다.
	// 그래서 최악의 경우 계산량은 약 81 * 30 * 9 * 30 정도이다.
	// 그러나 최악의 경우에도 계산량이 정해 져 있으므로 O(1) 로 실행 시간 복잡도를 표현할 수 있다.
	// 공간 복잡도도 O(1) 이다.
}

func solveSudoku(board [][]byte) {
	// 스도쿠를 재귀적으로 푸는 함수를 호출한다.
	solve(board, 0)
}

