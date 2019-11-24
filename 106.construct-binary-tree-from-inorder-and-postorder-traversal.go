/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (41.92%)
 * Likes:    1114
 * Dislikes: 25
 * Total Accepted:    178.7K
 * Total Submissions: 426.2K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * Given inorder and postorder traversal of a tree, construct the binary tree.
 * 
 * Note:
 * You may assume that duplicates do not exist in the tree.
 * 
 * For example, given
 * 
 * 
 * inorder = [9,3,15,20,7]
 * postorder = [9,15,7,20,3]
 * 
 * Return the following binary tree:
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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

func findIndex(val int, inorder []int) int {
	for i, n := range inorder {
		if n == val {
			return i
		}
	}
	return -1
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// 이 문제는 105 번의 preorder, inorder 문제의 변형으로 
	// postorder 는 제일 뒤 요소가 루트임을 이용해서 루트를 먼저 정하고 하위를
	// inorder 에서 루트를 기준으로 나누어 재귀호출하여 해결했다.
	// 이때 right 를 먼저 처리해 주는 것이 포인트이다.
	// 왜냐하면 postorder 이기 때문에 뒤에서 접근하면 오른쪽 요소가 먼저 postorder 에 나타나기 때문이다.

	root := TreeNode{Val: postorder[len(postorder)-1]}

	rootIndex := findIndex(root.Val, inorder)
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	nextPostorder:= len(postorder)-1
	if len(rightInorder) > 0 {
		root.Right = buildTree(rightInorder, postorder[:nextPostorder])
		nextPostorder -= len(rightInorder)
	}

	if len(leftInorder) > 0 {
		root.Left = buildTree(leftInorder, postorder[:nextPostorder])
	}


	return &root

}

// @lc code=end

