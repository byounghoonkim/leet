/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (44.97%)
 * Likes:    2879
 * Dislikes: 66
 * Total Accepted:    499.2K
 * Total Submissions: 1.1M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given a binary tree, check whether it is a mirror of itself (ie, symmetric
 * around its center).
 * 
 * For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 
 * 
 * But the following [1,2,2,null,3,null,3] is not:
 * 
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 
 * 
 * Note:
 * Bonus points if you could solve it both recursively and iteratively.
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

 func checkSubTree(l, r *TreeNode) bool {
	 if l == nil && r == nil {
		 return true
	 }

	 if l != nil && r != nil {
		 if l.Val != r.Val {
			 return false
		 }

		 return checkSubTree(l.Left, r.Right) && checkSubTree(l.Right, r.Left)
	 }

	 return false
 }

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkSubTree(root.Left, root.Right)
}
// @lc code=end

