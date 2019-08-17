/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
 *
 * https://leetcode.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (26.12%)
 * Likes:    2213
 * Dislikes: 332
 * Total Accepted:    446.9K
 * Total Submissions: 1.7M
 * Testcase Example:  '[2,1,3]'
 *
 * Given a binary tree, determine if it is a valid binary search tree (BST).
 *
 * Assume a BST is defined as follows:
 *
 *
 * The left subtree of a node contains only nodes with keys less than the
 * node's key.
 * The right subtree of a node contains only nodes with keys greater than the
 * node's key.
 * Both the left and right subtrees must also be binary search trees.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * Input: [2,1,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 *
 * Input: [5,1,4,null,null,3,6]
 * Output: false
 * Explanation: The root node's value is 5 but its right child's value is 4.
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

func isValid(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val <= root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val >= root.Right.Val {
		return false
	}
	if min != nil && min.Val >= root.Val {
		return false
	}
	if max != nil && max.Val <= root.Val {
		return false
	}

	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

