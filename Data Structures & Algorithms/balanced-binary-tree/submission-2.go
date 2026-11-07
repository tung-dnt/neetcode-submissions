/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	ok := true
	var dfs func(*TreeNode) int 
	dfs = func (n *TreeNode) int {
		if n == nil || !ok {
			return 0
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		if abs(l-r) > 1 {
			ok = false
		}
		return max(l,r) + 1
	}
	dfs(root)
	return ok
}

func abs(n int) int{
	if n < 0 {
		return -n
	}
	return n
}