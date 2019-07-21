/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 *
 * https://leetcode.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (44.55%)
 * Likes:    2375
 * Dislikes: 87
 * Total Accepted:    431.5K
 * Total Submissions: 968.5K
 * Testcase Example:  '2'
 *
 * You are climbing a stair case. It takes n steps to reach to the top.
 * 
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can
 * you climb to the top?
 * 
 * Note: Given n will be a positive integer.
 * 
 * Example 1:
 * 
 * 
 * Input: 2
 * Output: 2
 * Explanation: There are two ways to climb to the top.
 * 1. 1 step + 1 step
 * 2. 2 steps
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: 3
 * Output: 3
 * Explanation: There are three ways to climb to the top.
 * 1. 1 step + 1 step + 1 step
 * 2. 1 step + 2 steps
 * 3. 2 steps + 1 step
 * 
 * 
 */

func climb(n, cur int, cache map[int]int) int {
	// 현재 위치가 n 이면 계단을 다 올라 온 것이므로 
	// 방법 한가지가 있다고 리턴한다.
	if cur == n {
		return 1
	}

	// 현재 위치가 n 보다 크면 스텝이 목적지 모다 멀리 간 것이므로
	// 0을 리턴해서 가지수를 더하지 않는다.
	if cur > n {
		return 0
	}

	// Dynamic programming 을 적용한다.
	if r := cache[cur]; r > 0 {
		return r 
	}

	// 한계단 올라가는 것과 두 계단 올라 가는 것으로 
	// 재귀 호출 한 다음 결과를 더해서 리턴한다.
	r:= climb(n, cur+1, cache) + climb(n, cur+2, cache)

	// Dynamic programming 을 적용하여 캐시 한다.
	cache[cur] = r

	return r
}

func climbStairs(n int) int {
	cache:= make(map[int]int)
	return climb(n, 0, cache)
}

