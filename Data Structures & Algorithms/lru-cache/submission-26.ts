class ListNode {
    key;
    val;
    next;
    prev;

    constructor(key, val, next, prev) {
        this.key = key;
        this.val = val;
        this.next = next;
        this.prev = prev;
    }
}

class LRUCache {
    readonly capacity: number;
    readonly map = new Map<number, ListNode>();

    head: ListNode;
    tail: ListNode;

    /**
     * @param {number} capacity
     */
    constructor(capacity: number) {
        this.capacity = capacity;
    }

    /**
     * @param {number} key
     * @return {number}
     */
    get(key: number): number {
        const node = this.map.get(key);

        if (!node) return -1;

        if (this.map.size === 1 || !node.next) {
            return node.val;
        }

        const temp = node;

        if (!node.prev) {
            this.head = node.next;
            this.head.prev = null;
        } else {
            node.prev.next = temp.next;
            node.next.prev = temp.prev;
        }

        this.tail.next = node;
        node.prev = this.tail

        this.tail = this.tail.next;
        this.tail.next = null;

        return temp.val;
    }

    /**
     * @param {number} key
     * @param {number} value
     * @return {void}
     */
    put(key: number, value: number): void {
        const node = this.map.get(key);

        if (node) {
            node.val = value;
            this.map.set(key, node);

            if (this.map.size === 1 || !node.next) {
                return;
            }

            const temp = node;

            if (!node.prev) {
                this.head = node.next;
                this.head.prev = null;
            } else {
                node.prev.next = temp.next;
                node.next.prev = temp.prev;
            }

            this.tail.next = node;
            node.prev = this.tail

            this.tail = this.tail.next;
            this.tail.next = null;

            return;
        }

        const length: number = this.map.size;

        if (length === this.capacity) {
            const removedKey = this.head.key;

            this.head = this.head.next;
            if (this.capacity !== 1) this.head.prev = null;

            this.map.delete(removedKey);
        }

        if (!this.head) {
            this.head = new ListNode(key, value, null, null);

            this.map.set(key, this.head);
        } else if (!this.tail) {
            this.tail = new ListNode(key, value, null, this.head);
            this.head.next = this.tail;

            this.map.set(key, this.tail);
        } else {
            this.tail.next = new ListNode(key, value, null, this.tail);
            this.tail = this.tail.next;

            this.map.set(key, this.tail);
        }
    }
}
