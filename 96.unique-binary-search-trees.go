/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

func num(n int, c map[int]int) int {
	if n <= 1 {
		return 1 
	}

	if c[n] > 0 {
		return c[n]
	}

	ret := 0
	for i := 0; i < n ; i++ {
		ret +=  num(n-i-1, c) * num(i, c) 
	}

	c[n] = ret
	return ret
}

func numTrees(n int) int {
	cache := make(map[int]int)
	return num(n, cache)
}

