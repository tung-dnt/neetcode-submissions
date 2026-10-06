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
     * @return {number}
     */
    pairSum(head: ListNode | null): number {
        let [slow, fast] = [head, head];
        let stack = [];
        let max = 0;

        while (fast && fast.next) {
            stack.push(slow);

            slow = slow.next;
            fast = fast.next.next;
        }

        while (slow && stack.length > 0) {
            const sum = slow.val + stack.pop().val;

            if (sum > max) max = sum;

            slow = slow.next;
        }

        return max;
    }
}
