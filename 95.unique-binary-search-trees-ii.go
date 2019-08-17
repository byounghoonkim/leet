/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (36.56%)
 * Likes:    1387
 * Dislikes: 119
 * Total Accepted:    148K
 * Total Submissions: 404.7K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func gen(n int, b int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{nil}
	}
	if n == 1 {
		return []*TreeNode{{Val: b}}
	}

	result := []*TreeNode{}

	for i := 0; i < n; i++ {

		lefts := gen(i, b)
		rights := gen(n-1-i, b+i+1)

		for _, l := range lefts {
			for _, r := range rights {
				node := TreeNode{
					Val:   b + i,
					Left:  l,
					Right: r,
				}

				result = append(result, &node)
			}
		}
	}

	return result
}

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return []*TreeNode{}
	}

	return gen(n, 1)
}

