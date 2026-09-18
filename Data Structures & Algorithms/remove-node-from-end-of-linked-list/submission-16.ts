class Solution {
    /**
     * @param {ListNode} head
     * @param {number} n
     * @return {ListNode}
     */
    removeNthFromEnd(head: ListNode | null, n: number): ListNode {
        let length: number = 0;
        let current: ListNode = head;

        while (current) {
            length++;
            current = current.next;
        }

        const removeIdx: number = length - n;

        if (removeIdx === 0) return head.next;

        current = head;
        let currentIdx: number = 0;

        while (current) {
            if (currentIdx === removeIdx - 1) {
                current.next = current.next.next;
                break;
            }

            current = current.next;
            currentIdx++;
        }

        return head;
    }
}
