/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func (*TreeNode) int
	dfs = func (n *TreeNode) int {
		if n == nil {
			return 0
		}
		l:=dfs(n.Left)
		r:=dfs(n.Right)
		res = max(l+r ,res)
		return max(l,r) + 1
	}
	dfs(root)
	return res
}
