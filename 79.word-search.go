/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (31.81%)
 * Likes:    1934
 * Dislikes: 97
 * Total Accepted:    303.2K
 * Total Submissions: 953.1K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 * 
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 * 
 * Example:
 * 
 * 
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 * 
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 * 
 * 
 */


func search(board [][]byte, word string, i, j, d, pi,pj int) bool {

	fmt.Println(i,j,d, board[i][j])

	if board[i][j] != word[d] {
		return false
	}

	if d == len(word)-1 {
		return true
	}

	if i-1 >= 0 && i-1 != pi {
		if search(board, word, i-1, j, d+1, i,j) {
			return true
		}
	}
	if j + 1 < len(board[0]) && j+1 != pj {
		if search(board, word, i, j+1, d+1, i,j) {
			return true
		}
	}
	if i+1 < len(board) && i+1 != pi {
		if search(board, word, i+1, j, d+1, i,j) {
			return true
		}
	}
	if j-1 >= 0 && j-1 != pj {
		if search(board, word, i, j-1, d+1, i,j) {
			return true
		}
	}
	
	return false
}

func exist(board [][]byte, word string) bool {

	for i, b:= range board {
		for j, _:= range b {
			if search(board, word, i,j,0, i,j) {
				return true
			}
		}
	}

	return false
}

