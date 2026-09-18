/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]struct{})
	cur := head
	for cur != nil {
		if _, exists := seen[cur]; exists {
			return true
		}
		seen[cur] = struct{}{}
		cur = cur.Next
	}
	return false
}
