import "slices"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pStack, qStack := []int{}, []int{}
	dfs(p,&pStack)
	dfs(q,&qStack)
	return slices.Equal(pStack, qStack)
}

func dfs(n *TreeNode, stack *[]int) {
	if n == nil {
		*stack = append(*stack, -101)
		return
	}
	*stack = append(*stack, n.Val)
	dfs(n.Left, stack)
	dfs(n.Right, stack)
}