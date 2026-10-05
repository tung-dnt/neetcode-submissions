/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// nil -> 1 -> *(2) -> 3 -> *(4) -> 5 -> nil
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
//     if right == 1 || left == right {
// 		return head
// 	}
// 	dummy := &ListNode{Next: head}
// 	prevL, prevR, L := dummy,dummy,dummy
// 	cur := dummy.Next
// 	count := 1
// 	for cur != nil && count <= right {
// 		switch count {
// 			case left - 1:
// 				prevL = cur
// 			case left:
// 				L = cur
// 			case right - 1:
// 				prevR = cur
// 			case right:
// 				tmp := cur.Next
				
// 				prevL.Next = cur
// 				prevR.Next = L // wrong
// 				cur.Next = L.Next
// 				L.Next = tmp
// 		}
// 		cur = cur.Next
// 		count++
// 	}
// 	return dummy.Next
// }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{Next: head}
    prev, cur := dummy, head // prev ends just before left; cur = node at left (stays as segment tail)

    for count := 1; count < right; count++ {
        if count < left { // phase 1: walk to left
            prev, cur = cur, cur.Next
            continue
        }
        // phase 2: move cur.Next to the front of the segment
        nxt := cur.Next
        cur.Next = nxt.Next
        nxt.Next = prev.Next
        prev.Next = nxt
    }
    return dummy.Next
}