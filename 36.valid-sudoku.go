import (
	"strconv"
)

/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 *
 * https://leetcode.com/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (43.41%)
 * Likes:    861
 * Dislikes: 317
 * Total Accepted:    242.4K
 * Total Submissions: 558.3K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be
 * validated according to the following rules:
 *
 *
 * Each row must contain the digits 1-9 without repetition.
 * Each column must contain the digits 1-9 without repetition.
 * Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without
 * repetition.
 *
 *
 *
 * A partially filled sudoku which is valid.
 *
 * The Sudoku board could be partially filled, where empty cells are filled
 * with the character '.'.
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * ⁠ ["5","3",".",".","7",".",".",".","."],
 * ⁠ ["6",".",".","1","9","5",".",".","."],
 * ⁠ [".","9","8",".",".",".",".","6","."],
 * ⁠ ["8",".",".",".","6",".",".",".","3"],
 * ⁠ ["4",".",".","8",".","3",".",".","1"],
 * ⁠ ["7",".",".",".","2",".",".",".","6"],
 * ⁠ [".","6",".",".",".",".","2","8","."],
 * ⁠ [".",".",".","4","1","9",".",".","5"],
 * ⁠ [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * [
 * ["8","3",".",".","7",".",".",".","."],
 * ["6",".",".","1","9","5",".",".","."],
 * [".","9","8",".",".",".",".","6","."],
 * ["8",".",".",".","6",".",".",".","3"],
 * ["4",".",".","8",".","3",".",".","1"],
 * ["7",".",".",".","2",".",".",".","6"],
 * [".","6",".",".",".",".","2","8","."],
 * [".",".",".","4","1","9",".",".","5"],
 * [".",".",".",".","8",".",".","7","9"]
 * ]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner
 * being
 * ⁠   modified to 8. Since there are two 8's in the top left 3x3 sub-box, it
 * is invalid.
 *
 *
 * Note:
 *
 *
 * A Sudoku board (partially filled) could be valid but is not necessarily
 * solvable.
 * Only the filled cells need to be validated according to the mentioned
 * rules.
 * The given board contain only digits 1-9 and the character '.'.
 * The given board size is always 9x9.
 *
 *
 */

func check(checker *[9][9]bool, i int, n byte) bool {
	// 이 함수는 9*9 2차원 배열을 포인터로 받아
	// i 행에 n 열에 이미 true 로 설정되어 있으면 false 를 리턴하고
	// false이면 true로 설정한 후 true를 리턴한다.
	num, error := strconv.Atoi(string(n))
	if error != nil {
		return true
	}

	num -= 1

	if checker[i][num] == true {
		return false
	}

	checker[i][num] = true
	return true
}

func isValidSudoku(board [][]byte) bool {

	// 스도쿠 각 행, 열, 작은 박스를 체크 하기 위한
	// 배열을 9*9 * 3개 준비한다.
	rchecker := [9][9]bool{}
	cchecker := [9][9]bool{}
	schecker := [9][9]bool{}

	for r, a := range board {
		for c, n := range a {
			if n == '.' {
				continue
			}

			// 행의 값을 검사
			if ret := check(&rchecker, r, n); ret == false {
				return ret
			}

			// 열의 값을 검사
			if ret := check(&cchecker, c, n); ret == false {
				return ret
			}

			// 작은 박스의 인덱스를 구하고
			sbox_index := (c / 3) + (r/3)*3
			// 구한 인덱스로 작은 박스의 유효성을 검사
			if ret := check(&schecker, sbox_index, n); ret == false {
				return ret
			}

			// 각 요소 검사 시 false가 한번이라도 나오면 전체 루프를 종료하고
			// false 를 리턴한다.
		}
	}

	return true

	// 이 알고리즘은 board 요소의 갯수 만큼 루프를 1회 실행 하므로
	// 실행 시간 복잡도는 O(n) 이다
	// 체커를 만들기 위해 board 요소 개수의 3배 만큼 메모리를 할당 했으므로
	// 3n 으로 공간 복잡도도 O(n) 이다.

}

