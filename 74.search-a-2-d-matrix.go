/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (35.04%)
 * Likes:    907
 * Dislikes: 108
 * Total Accepted:    234.7K
 * Total Submissions: 669.9K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 * 
 * 
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 * 
 */
func searchMatrix(matrix [][]int, target int) bool {
	// Bianry Search Version
	l := 0
	r := len(matrix)-1

	row := -1
	for l <= r {
		if len(matrix[l]) == 0 {
			l++ 
			continue
		}
		if matrix[l][0] > target {
			row = l -1 
			break
		}
		if len(matrix[r]) == 0 {
			r --
			continue
		}

		if matrix[r][0] <= target {
			row = r 
			break
		}

		if l + 1 == r {
			row = l
			break
		}

		m := (l+r)/2
		if matrix[m][0] > target {
			r = m
		} else {
			l = m
		}
	}

	if row == -1 {
		return false
	}

	l = 0 
	r = len(matrix[row]) - 1

	if target > matrix[row][r] {
		return false
	}

	for ;; {
		m := (l+r)/2
		if matrix[row][l] == target || 
			matrix[row][r] == target || 
			matrix[row][m] == target {
				return true
			}
		
		if l == m || r == m {
			return false
		}

		if matrix[row][m] > target {
			r = m
		} else {
			l = m
		}
	}

	return false

}
func searchMatrix2(matrix [][]int, target int) bool {

	// Simple Search Version

	row := -1
	for i := 0 ; i < len(matrix) ; i ++ {
		if len(matrix[i]) == 0 {
			continue
		}

		if target < matrix[i][0] {
			row = i-1
			break
		}
		if i == len(matrix)-1 {
			row = len(matrix) -1 
			break
		}
	}

	if row == -1 {
		return false
	}

	for _, a := range(matrix[row]) {
		if target == a {
			return true
		}
	}

	return false
}

