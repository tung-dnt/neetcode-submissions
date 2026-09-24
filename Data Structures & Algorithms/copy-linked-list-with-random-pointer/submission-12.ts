// class Node {
//   constructor(val, next = null, random = null) {
//       this.val = val;
//       this.next = next;
//       this.random = random;
//   }
// }

class Solution {
    /**
     * @param {Node} head
     * @return {Node}
     */
    copyRandomList(head: Node | null): Node {
        if (head === null) return null;

        const map = new Map();
        const newHead = new Node(head.val, null, head.random);
        map.set(head, newHead);

        let current = head;
        let newCurrent = newHead;
        while (current.next !== null) {
            newCurrent.next = new Node(current.next.val, null, current.next.random);
            map.set(current.next, newCurrent.next);

            newCurrent = newCurrent.next;
            current = current.next;
        }

        newCurrent = newHead;
        while (newCurrent !== null) {
            newCurrent.random = map.get(newCurrent.random);

            newCurrent = newCurrent.next;
        }

        return newHead;
    }
}
