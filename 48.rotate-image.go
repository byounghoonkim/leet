/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

func trans(i, j, n int) (int, int) {
	// 90도 회전환 좌표를 리턴한다.
	return j, n - i
}

func swap(matrix [][]int, i, j, a, b int) {
	matrix[i][j], matrix[a][b] = matrix[a][b], matrix[i][j]
}

func rotate(matrix [][]int) {

	l := len(matrix)
	n := l - 1

	// i는 절반 만큼만 루프를 돈다. 이유는 나머지 부분은 회적 시에 이미 처리 되기 때문이다
	for i := 0; i < l/2; i++ {
		// j 는 i 부터 시작 해서 l - i 에서 한 줄이 더 뺀 부분 까지 루프를 돈다.
		// i 부터 시작 하는 이유는 i 값 만큼 앞 부분은 이미 회전을 완료한 상태이므로 건드리면 안된다.
		// l-i 에서 1을 더 배는 이유는 마지막 줄은 첫 요소가 회전 할 때 이미 회전 된 요소기기 때문이다.
		for j := i; j < l-1-i; j++ {
			a, b := i, j
			// 3번 좌표를 변환해 가며 값을 스왑한다.
			for k := 0; k < 3; k++ {
				a, b = trans(a, b, n)
				swap(matrix, i, j, a, b)
			}
		}
	}

	// 이 알고리즘은 n/2 * (n-i) * 3 정도의 실행 시간을 가진다.
	// O(n^2) 정도의 실행 시간 복잡도를 가진다.
}

