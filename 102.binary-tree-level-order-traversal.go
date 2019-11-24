/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (50.99%)
 * Likes:    1966
 * Dislikes: 54
 * Total Accepted:    467.3K
 * Total Submissions: 916.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
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
 * return its level order traversal as:
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
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

 func printNodes(list []*TreeNode, results *[][]int) []*TreeNode {

	ret := []*TreeNode{}

	result := []int {}
	for _, n := range list {
		result = append(result, n.Val)
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

func levelOrder(root *TreeNode) [][]int {
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

