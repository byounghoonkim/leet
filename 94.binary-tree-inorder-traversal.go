/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 *
 * https://leetcode.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (57.85%)
 * Likes:    1820
 * Dislikes: 77
 * Total Accepted:    503.8K
 * Total Submissions: 870.9K
 * Testcase Example:  '[1,null,2,3]'
 *
 * Given a binary tree, return the inorder traversal of its nodes' values.
 *
 * Example:
 *
 *
 * Input: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * Output: [1,3,2]
 *
 * Follow up: Recursive solution is trivial, could you do it iteratively?
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

func traverse(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	traverse(node.Left, result)
	*result = append(*result, node.Val)
	traverse(node.Right, result)
}

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	traverse(root, &ret)
	return ret
}

