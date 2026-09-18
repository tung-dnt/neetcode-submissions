class Solution {
    /**
     * @param {ListNode} head
     * @param {number} n
     * @return {ListNode}
     */
    removeNthFromEnd(head: ListNode | null, n: number): ListNode {
        let length: number = 0
        let current: ListNode = head

        while (current) {
            length++
            current = current.next
        }

        const removeIdx: number = length - n

        if (removeIdx === 0) return length === 1 ? null : head.next

        let newCurrent = head
        let currentIdx: number = 0

        while (newCurrent) {
            if (currentIdx === removeIdx - 1) {                
                newCurrent.next = newCurrent.next.next
                break
            }

            newCurrent = newCurrent.next
            currentIdx++
        }

        return head
    }
}
