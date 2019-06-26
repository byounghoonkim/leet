import "fmt"

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
	if level == 81 {
		return true
	}

	r := level / 9
	c := level % 9

	if '.' != board[r][c] {
		return solve(board, level+1)
	}

	candi := candidates(board, r, c)

	if 0 == len(candi) {
		fmt.Println(r, c)
	}

	for _, ca := range candi {
		board[r][c] = i2b(ca)

		if true == solve(board, level+1) {
			return true
		}
	}
	return false
}

func solveSudoku(board [][]byte) {

	solve(board, 0)

}

