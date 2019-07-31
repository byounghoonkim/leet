/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
 */

func height(matrix[][]byte, i, j int) int {
	for a:= j; a < len(matrix[i]); a++ {
		if matrix[i][a] != '1' {
			return a-j
		}
	}
	return len(matrix[i]) - j
}

func min(a, b int) int {
	if a > b { return b}
	return a
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

// 이문제는 직전 문제인 84번 솔루션을 응용했다.
// 알고리즘의 아이디어는 각 열을 x 축으로 둔
// 가장 큰 사각형을 구혀는 문제로 변형하였다.
func maximalRectangle(matrix [][]byte) int {
	ret := 0

	for i:= 0 ; i < len(matrix); i++ {
		for k := 0 ; k < len(matrix[0]); k ++ {
			min_h := height(matrix, i, k) 
			for j := i ; j < len(matrix); j++ {
				min_h = min(min_h, height(matrix, j, k))

				rect := (j-i+1)*min_h
				ret = max(ret, rect)
			}
		}
	}


	return ret
    
}

