/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 2 pointer

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	// Nếu l start ở head thì khi dừng sẽ dừng đúng ngay vị trí cần replace
	// Muốn replace thì phải nối Next của node trước vào node sau, nên phải lùi left về 1 node trước khi duyệt
	// -> Tạo dummy node
	l,r := dummy, head
	for n != 0 {
		r = r.Next
		n--
	}
	for r != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
