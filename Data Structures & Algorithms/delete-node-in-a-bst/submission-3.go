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
		return nil
	}

	if val > root.Val {
		root.Right = deleteNode(root.Right, val)
	} else if val < root.Val {
		root.Left = deleteNode(root.Left, val)
	} else {
		// Trường hợp 1 & 2: Node không có con bên trái, hoặc là node lá (cả 2 con đều nil)
		// Nếu Left là nil, ta chỉ cần trả về Right (dù Right có là nil hay không thì kết quả vẫn đúng).
		if root.Left == nil {
			return root.Right
		}

		// Trường hợp 3: Node không có con bên phải
		if root.Right == nil {
			return root.Left
		}

		// Trường hợp 4: Node có cả 2 con
		// Lúc này ta không cần kiểm tra != nil nữa vì nó đã vượt qua 2 câu lệnh if ở trên
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		cur.Left = root.Left
		return root.Right
	}

	return root
}
