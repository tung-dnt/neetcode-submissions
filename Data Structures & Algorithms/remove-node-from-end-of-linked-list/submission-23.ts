class Solution {
    /**
     * @param {ListNode} head
     * @param {number} n
     * @return {ListNode}
     */
    removeNthFromEnd(head: ListNode | null, n: number): ListNode {
        const dummy = new ListNode(null, head)

        let left = dummy
        let right = head

        while (n > 0) {
            right = right.next
            n--
        }

        while (right !== null) {
            left = left.next
            right = right.next
        }

        left.next = left.next.next

        return dummy.next
    }
}
