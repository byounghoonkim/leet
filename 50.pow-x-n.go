/*
 * @lc app=leetcode id=50 lang=golang
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (28.13%)
 * Likes:    871
 * Dislikes: 2140
 * Total Accepted:    330.8K
 * Total Submissions: 1.2M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (x^n).
 *
 * Example 1:
 *
 *
 * Input: 2.00000, 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: 2.10000, 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: 2.00000, -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 * Note:
 *
 *
 * -100.0 < x < 100.0
 * n is a 32-bit signed integer, within the range [−2^31, 2^31 − 1]
 *
 *
 */
func myPow(x float64, n int) float64 {
	ret := 1.0
	target := n
	newx := x
	if n < 0 {
		newx = 1 / x
		target = n * -1
	}

	for target != 0 {
		remain := target
		r := 1.0
		// i를 2배수씩 늘려가며 target 이하까지 루프를 돈다
		// 이때 값은 결과의 제곱만큼 계산해 나간다.
		// 이렇게 접근 해야 n 이 아주 크거나 음수로 아주 작을 때
		// TLE 가 발생하지 않는다.
		for i := 1; i <= target; i *= 2 {
			if i == 1 {
				r = newx
			} else {
				r *= r
			}
			remain = target - i
		}

		// 중간 결과 값을 최종 값에 반영한다.
		ret *= r

		// 나머지 지수를 target 으로 지정하고
		// 다시 위와 같은 로직을 반복해 target 이 0이 될때까지
		// 반복한다.
		target = remain
	}

	return ret

	// 단순히 접근 하면 x 를 n 번 곱해야 하므로
	// O(n)이 된다.
	// 그런데 n 이 아주 높을 때 수행 시간이 너무 높아 지기 때문에
	// 이 알고리즘은 결과의 제곱하는 방법으로 지수를 2 의 배수 만큼 씩
	// 줄이는 방법을 택했다.
	// O(logn) 정도의 수행 시간 복잡도를 가진다.
	// TODO 이 부분은 조금만 더 정교히 계산해 보자.
}
