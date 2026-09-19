/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	cur := head
	for cur != nil {
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	removeIdx := len(nodes) - n
	if removeIdx == 0 {
		return head.Next
	}
	nodes[removeIdx-1].Next = nodes[removeIdx].Next
	return head
}