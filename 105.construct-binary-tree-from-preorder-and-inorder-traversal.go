/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (44.25%)
 * Likes:    2270
 * Dislikes: 61
 * Total Accepted:    275.1K
 * Total Submissions: 621.5K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * Given preorder and inorder traversal of a tree, construct the binary tree.
 * 
 * Note:
 * You may assume that duplicates do not exist in the tree.
 * 
 * For example, given
 * 
 * 
 * preorder = [3,9,20,15,7]
 * inorder = [9,3,15,20,7]
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	// preorder 에서는 항상 첫번째 요소가 루트이다.
	root := TreeNode{Val: preorder[0]}

	// root 를 inorder 리스트에서 찾아서 root 를 기준으로 두 리스트로 나눈다.
	rootIndex := findIndex(root.Val, inorder)
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	nextPreorder := 1
	if len(leftInorder) > 0 {
		// 왼편에 요소가 있다면 재귀 호출로 서브 트리를 생성한다. 
		// 이때 preorder 는 루트 요소를 제외한 1: 까지로 정한다.
		root.Left = buildTree(preorder[nextPreorder:], leftInorder)  


		// 이 문제의 가장 어려운 부분이다. 다음 preorder 리스트를 위해 left inorder 의 개수 만큼
		// 더해서 사용한다.
		nextPreorder += len(leftInorder)
	}

	if len(rightInorder) > 0 {
		root.Right = buildTree(preorder[nextPreorder:], rightInorder)  
	}

	return &root

}
// @lc code=end

