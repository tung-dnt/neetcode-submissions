/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val < val {
		root.Right = deleteNode(root.Right, val)
	} else if root.Val > val {
		root.Left = deleteNode(root.Left, val)
	} else {
		if root.Left != nil && root.Right != nil {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else {
			return nil
		}
	}

	return root
}
