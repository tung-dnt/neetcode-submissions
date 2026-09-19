/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import "slices"

// Array converting
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	arr:=make([]int, 0)
	cur:=head
	for cur != nil  {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	idxToDelete := len(arr)-n
	arr = slices.Delete(arr, idxToDelete, idxToDelete+1)

	dummyNode := &ListNode{}
	curNode := dummyNode
	for _, val := range arr {
		newNode := &ListNode{Val: val}
		curNode.Next = newNode
		curNode = newNode
	}

	return dummyNode.Next
}