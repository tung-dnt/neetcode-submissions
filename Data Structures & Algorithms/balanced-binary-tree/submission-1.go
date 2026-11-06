/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return checkHeight(root) != -1
}

// checkHeight returns the height of the subtree if balanced, or -1 if unbalanced.
func checkHeight(node *TreeNode) int {
    if node == nil {
        return 0
    }

    leftHeight := checkHeight(node.Left)
    if leftHeight == -1 {
        return -1
    }

    rightHeight := checkHeight(node.Right)
    if rightHeight == -1 {
        return -1
    }

    diff := leftHeight - rightHeight
    if diff < -1 || diff > 1 {
        return -1
    }

    if leftHeight > rightHeight {
        return leftHeight + 1
    }
    return rightHeight + 1
}