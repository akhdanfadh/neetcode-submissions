/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil { return true }
	if root == nil { return false }
	if sameTree(root, subRoot) { return true }
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(p, q *TreeNode) bool {
	if p == nil && q == nil { return true }
	if p != nil && q != nil && p.Val == q.Val {
		return sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
	}
	return false
}