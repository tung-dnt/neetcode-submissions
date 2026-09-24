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
        map.set(head, new Node(head.val, null, null));

        let current = head;
        while (current.next !== null) {
            const newNode = new Node(current.next.val, null, null)

            map.get(current).next = newNode;
            map.set(current.next, newNode);
            current = current.next;
        }

        current = head;
        while (current !== null) {
            map.get(current).random = map.get(current.random);

            current = current.next;
        }

        return map.get(head);
    }
}
