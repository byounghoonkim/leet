/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
 *
 * https://leetcode.com/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Hard (35.40%)
 * Likes:    876
 * Dislikes: 48
 * Total Accepted:    125.1K
 * Total Submissions: 353.5K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * Two elements of a binary search tree (BST) are swapped by mistake.
 *
 * Recover the tree without changing its structure.
 *
 * Example 1:
 *
 *
 * Input: [1,3,null,null,2]
 *
 * 1
 * /
 * 3
 * \
 * 2
 *
 * Output: [3,1,null,null,2]
 *
 * 3
 * /
 * 1
 * \
 * 2
 *
 *
 * Example 2:
 *
 *
 * Input: [3,1,4,null,null,2]
 *
 * ⁠ 3
 * ⁠/ \
 * 1   4
 * /
 * 2
 *
 * Output: [2,1,4,null,null,3]
 *
 * ⁠ 2
 * ⁠/ \
 * 1   4
 * /
 * ⁠ 3
 *
 *
 * Follow up:
 *
 *
 * A solution using O(n) space is pretty straight forward.
 * Could you devise a constant space solution?
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

func recover(root *TreeNode, min, max *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if min != nil && min.Val >= root.Val {
		return root
	}
	if max != nil && max.Val <= root.Val {
		return root
	}

	l := recover(root.Left, min, root)
	r := recover(root.Right, root, max)

	if l != nil && r != nil {
		l.Val, r.Val = r.Val, l.Val
	} else {
		if l != nil {
			if l.Val > root.Val {
				l.Val, root.Val = root.Val, l.Val
			} else {
				return l
			}
		} else if r != nil {
			if r.Val < root.Val {
				r.Val, root.Val = root.Val, r.Val
			} else {
				return r
			}
		} else {
			// ?
		}
	}
	return nil
}

func recoverTree(root *TreeNode) {
	recover(root, nil, nil)
}

