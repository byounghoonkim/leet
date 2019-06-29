/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

func helper(a []int, c int, j int, cache []int) int {
	// 현재 인덱스 c 가 마지막에 도달 했다면 0을 리턴한다.
	if len(a)-1 == c {
		return 0
	}

	// cache 에 현재 인덱스 c 로 부터 마지막 까지의 점프 최소 값이 계산 되어 있다면
	// 그것을 사용해 바로 리턴해 중복 계산을 줄인다.
	if cache[c] > 0 {
		return cache[c]
	}

	// 최소값 기본 값은 a 의 길이 (최대값) 으로 지정한다.
	min := len(a)

	// a[c] 값 만큼 루프를 돌면서 점프한 인덱스로 재귀 호출 한다.
	for i := 1; i <= a[c]; i++ {
		if c+i >= len(a) {
			break
		}

		m := helper(a, c+i, j+1, cache) + 1
		if min > m {
			min = m
		}
	}

	// 최소 값을 현재 인덱스 c 에 캐시 한다.
	cache[c] = min
	return min
}

func jump(nums []int) int {

	cache := make([]int, len(nums))
	return helper(nums, 0, 0, cache)

	// 이 알고리즘은 cache 없이 탐색 할때
	// 최악의 경우 5 5 5 5 5 5 와 같이 충분히 점프 할 수 있는 숫자의 배열일 경우
	// 5 * 4 * 3 ... * 1 의 루프를 돌게 되고 O(n!) 을 가진다.
	// 그러나 동적 프로그래밍으로 중복 계산을 제거해 time limite exceed 를 회피 한다.
	// TODO 동적 프로그래밍 적용 시 Big O 표기법을 계산해 보자!
}

