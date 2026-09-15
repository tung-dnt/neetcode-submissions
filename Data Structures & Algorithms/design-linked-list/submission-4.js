class ListNode {
    constructor(val, prev = null, next = null) {
        this.prev = prev;
        this.val = val;
        this.next = next;
    }
}

class MyLinkedList {
    constructor() {
        this.dummyHead = new ListNode(0);
        this.dummyTail = new ListNode(0);
        this.dummyHead.next = this.dummyTail;
        this.dummyTail.prev = this.dummyHead;
        this.size = 0;
    }
    /**
     * @param {number} index
     * @return {ListNode}
     */
    getPrev(index) {
        let curr = this.dummyHead
        for(let i = 0; i < index; i++){
            curr = curr.next
        }
        return curr
    }

    /**
     * @param {number} index
     * @return {number}
     */
    get(index) {
        if(index < 0 || index >= this.size) return -1
        
        return this.getPrev(index).next.val
    }

    /**
     * @param {number} val
     * @return {void}
     */
    addAtHead(val) {
        this.addAtIndex(0,val)
    }

    /**
     * @param {number} val
     * @return {void}
     */
    addAtTail(val) {
        this.addAtIndex(this.size, val)
    }

    /**
     * @param {number} index
     * @param {number} val
     * @return {void}
     */
    addAtIndex(index, val) {
        if(index < 0 || index > this.size) return
        let currHead = this.getPrev(index).next;
        let prevHead = currHead.prev
        const newNode = new ListNode(val, prevHead, currHead)
        currHead.prev = newNode
        prevHead.next = newNode
        this.size++
    }

    /**
     * @param {number} index
     * @return {void}
     */
    deleteAtIndex(index) {
        if(index < 0 || index >= this.size) return
        let curr = this.getPrev(index).next;
        let prev = curr.prev
        let next = curr.next
       prev.next = next
        next.prev = prev
        this.size--
    }
}
