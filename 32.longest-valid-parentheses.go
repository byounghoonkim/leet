/*
 * @lc app=leetcode id=32 lang=golang
 *
 * [32] Longest Valid Parentheses
 */

func ptoi(c rune) int {
	if c == ')' {
		return -1
	}
	if c == '(' {
		return 1
	}
	return 0
}

func longestValidParentheses(s string) int {

	// return value
	ret := 0

	// 괄호의 현재 상태를 나타내는 정수
	// ==0: 유효 상태
	// >0 : 유효 가능 상태
	// <0 : 유효 불가( 가 더 많이 발생
	state := 0

	for i, _ := range s {
		for j := i; j < len(s); j++ {
			state += ptoi(rune(s[j]))

			// 현재 상태가 0 이면 유효 한 상태 이므로
			// 현재 길이를 구해 리턴할 값 보다 크다면 리턴 값을 대체한다.
			if state == 0 {
				length := j - i + 1
				if length > ret {
					ret = length
				}
			}

			// 상태가 -1 이라면 앞으로도 유효할 가능성이 없으므로
			// 루프를 멈추고 i 를 앞으로 진행할 수 있도록 한다.
			if state == -1 {
				break
			}
		}

		// i 가 진행 하기 전에 기존의 state 값은 0으로 초기화 한다.
		state = 0
	}

	return ret

	// 이 알고리즘은 s의 길이 n 에 대해 n * (n-1) * ... * 1 만큼 수행 해야 하므로
	// 시간 복잡도는 O(N!) 이다.
	// 공간 복잡도는 로컬 변수 이외에 추가로 할당하지 않으므로 O(1)이다.
}

