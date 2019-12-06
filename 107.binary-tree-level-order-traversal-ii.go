/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (49.06%)
 * Likes:    920
 * Dislikes: 173
 * Total Accepted:    262.6K
 * Total Submissions: 535.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the bottom-up level order traversal of its
 * nodes' values. (ie, from left to right, level by level from leaf to root).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its bottom-up level order traversal as:
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printNodes(list []*TreeNode, results *[][]int) []*TreeNode {

	ret := []*TreeNode{}

	result := []int{}
	for _, n := range list {
		result = append(result, n.Val)
		if n.Left != nil {
			ret = append(ret, n.Left)
		}
		if n.Right != nil {
			ret = append(ret, n.Right)
		}
	}

	if len(result) > 0 {
		// 이 문제는 102번 문제의 풀이에서
		// 아래 결과 슬라이스를 앞에서 부터 추가하도록 수정하여
		// 해답을 만들었다.
		*results = append([][]int{result}, *results...)
	}

	return ret
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	list := []*TreeNode{root}

	for len(list) > 0 {
		list = printNodes(list, &ret)
	}

	return ret

}

// @lc code=end

