/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }https://imagedelivery.net/CLfkmk9Wzy8_9HRyug4EVA/f1a84528-91c2-4ac2-8b0f-ed6b57881400/public$0
 * }
 */
class Solution {
    /**
     * @param {ListNode} head
     * @param {number} left
     * @param {number} right
     * @return {ListNode}
     */
    reverseBetween(head: ListNode | null, left: number, right: number): ListNode {
        let l = null;
        let r = null;
        let prev = null;
        let current = head;
        let count = 1;

        while (current !== null) {
            if (count === left - 1) {
                l = current;
            } else if (count === right + 1) {
                r.next = current;

                break;
            }

            if (count >= left && count <= right) {
                if (count === left) r = current;
                else if (count === right) {
                    if (l) l.next = current;
                    else head = current;
                }

                let temp = current.next;
                current.next = prev;
                prev = current;

                current = temp;
            } else {
                // prev = current;
                current = current.next;
            }

            count++;
        }

        return head;
    }
}
