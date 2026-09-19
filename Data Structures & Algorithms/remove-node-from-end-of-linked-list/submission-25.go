/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Count elements, traverse back to identify removed node - 12min
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLen := 0
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	removeIdx := listLen - n
	if removeIdx == 0 {
		return head.Next
	}
	i := 0
	dummyNode := &ListNode{Next: head}
	cur = dummyNode.Next
	for cur != nil {
		if i ==removeIdx - 1 {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
		i++
	}

	return dummyNode.Next
}