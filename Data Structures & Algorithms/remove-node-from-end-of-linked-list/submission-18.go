/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Count elements, traverse back to identify removed node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listLen := 0
	cur := head
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	idxToDelete := listLen - n
	i := 0
	dummyNode := &ListNode{Next: head}
	cur = dummyNode.Next
	for cur != nil {
		if i ==idxToDelete - 1 {
			cur.Next = cur.Next.Next
			break
		} else if i == idxToDelete {
			dummyNode.Next = dummyNode.Next.Next
			break
		}
		cur = cur.Next
		i++
	}

	return dummyNode.Next
}