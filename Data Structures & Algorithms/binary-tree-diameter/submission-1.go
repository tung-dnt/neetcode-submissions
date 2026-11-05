/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := -1
	var dfs func(*TreeNode) int
	dfs = func (node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		res = max(res, l+r)
		return max(l,r) + 1
	}
	dfs(root)
	return res
}
