/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (33.46%)
 * Likes:    883
 * Dislikes: 231
 * Total Accepted:    142.3K
 * Total Submissions: 425.3K
 * Testcase Example:  '3\n3'
 *
 * The set [1,2,3,...,n] contains a total of n! unique permutations.
 * 
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 * 
 * 
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 * 
 * 
 * Given n and k, return the k^th permutation sequence.
 * 
 * Note:
 * 
 * 
 * Given n will be between 1 and 9 inclusive.
 * Given k will be between 1 and n! inclusive.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: n = 3, k = 3
 * Output: "213"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 4, k = 9
 * Output: "2314"
 * 
 * 
 */

func factorial(x int) int {
	if x == 1 {
		return 1 
	}
	return x * factorial(x-1)
}

func getPermutation(n int, k int) string {
	ret := ""

	// 조합 가능한 수를 슬라이스로 만든다. 
	c:= make([]int, n)
	for i, _ := range(c) {
		c[i] = i+1
	}

	ck := k

	// 각 자리수 만큼 루프를 돈다.
	for i := 0; i < n ; i++ {
		// k가 첫번째 이면 현재 조합 가능한 슬라이스를 그대로
		// 결과에 넣고 리턴한다.
		if ck == 1 {
			for _, a := range(c) {
				ret += strconv.Itoa(a)
			}
			break
		}

		// 현재 자리를 제외한 나머지 수들의 조합 가능 수(팩토리얼)
		// 을 k - 1 로 나누면 현재 자리수에 들어갈 수의 
		// 인덱스를 알 수 있다.
		j:= (ck-1)/factorial(n-i-1)

		// j 의 수를 결과에 추가 한다.
		ret += strconv.Itoa(c[j])

		// c에서 j 번째 자리를 제외하고 나머지 조합 가능한 슬라이스를
		// 다시 만든다
		c = append(c[:j], c[j+1:]...)

		// ck에서 조합 가능한 수의 나머지 연산을 통해 
		// 하위 순열에서 몇번째 순열인지 계산해 놓는다.
		ck = ck%factorial(n-i-1)

		// ck 가 0 이면 마지막 순열이므로 다시 마지막 번째로 돌려 놓는다.
		if ck == 0 {
			ck = factorial(n-i-1)
		}

	}

	return ret

	// TODO 수행 시간 복잡도를 계산해 보자.

}

