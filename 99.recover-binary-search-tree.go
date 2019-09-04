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

 func travel(root *TreeNode, arr *[]*TreeNode) {
	 if root.Left == nil && root.Right == nil {
		*arr = append(*arr, root)
		return
	 }

	 if root.Left != nil {
		 travel(root.Left, arr)
	 }

	*arr = append(*arr, root)
	 if root.Right != nil {
		 travel(root.Right, arr)
	 }

 }
func recoverTree(root *TreeNode) {
	arr := []*TreeNode{}
	travel(root, &arr)

	s1 := -1
	s2 := -1
	for i:= 0; i < len(arr); i++ {
		fmt.Println(arr[i].Val)
		if i > 0 {
			if arr[i].Val < arr[i-1].Val {
				if s1 == -1 {
					s1 = i
				} else {
					s2 = i
				}
			}
		}
	}

	if s1 != -1 {
		if s2 == -1 {
			arr[s1-1].Val, arr[s1].Val = arr[s1].Val, arr[s1-1].Val
		} else {
			arr[s1-1].Val, arr[s2].Val = arr[s2].Val, arr[s1-1].Val
		}
	}
}

