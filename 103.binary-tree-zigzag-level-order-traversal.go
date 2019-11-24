/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (44.10%)
 * Likes:    1342
 * Dislikes: 74
 * Total Accepted:    275.4K
 * Total Submissions: 624.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the zigzag level order traversal of its nodes'
 * values. (ie, from left to right, then right to left for the next level and
 * alternate between).
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
 * return its zigzag level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
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

 func printNodes(list []*TreeNode, results *[][]int, forward bool) []*TreeNode {

	ret := []*TreeNode{}

	result := []int {}
	for _, n := range list {

		if forward {
			result = append(result, n.Val)
		} else {
			result = append([]int{n.Val}, result...)
		}

		if n.Left != nil {
			ret = append(ret, n.Left)
		}
		if n.Right!= nil {
			ret = append(ret, n.Right)
		}
	}

	if len(result) > 0 {
		*results = append(*results, result)
	}

	return ret
 }

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}

	list := []*TreeNode{root}

	forward := true
	for len(list) > 0 {
		list = printNodes(list, &ret, forward)
		forward = !forward
	}

	return ret

}
// @lc code=end

