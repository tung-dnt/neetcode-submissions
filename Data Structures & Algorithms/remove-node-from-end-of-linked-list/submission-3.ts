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
        let current2: ListNode = new ListNode(null, head)
        let giveMeAHead = current2
        let currentIdx: number = 0

        while (current2) {
            if (currentIdx === removeIdx) {                
                current2.next = current2.next.next
                break
            }

            current2 = current2.next
            currentIdx++
        }

        return giveMeAHead.next
    }
}
