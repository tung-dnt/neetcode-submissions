/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

class Solution {
    /**
     * @param {ListNode} head
     * @return {ListNode}
     */
    reverseList(head: ListNode | null): ListNode {
        if (!head) return null;

        let current = head;
        let newHead: ListNode = null;

        while (current !== null) {
            const newNode = new ListNode(current.val, newHead);

            newHead = newNode

            current = current.next
        }

        return newHead;
    }
}
