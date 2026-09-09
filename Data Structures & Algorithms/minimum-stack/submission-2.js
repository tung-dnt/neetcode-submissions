class MinStack {
    stack = [];
    minStack = []

    constructor() {}

    /**
     * @param {number} val
     * @return {void}
     */
    push(val) {
        if(this.minStack.length === 0) {this.minStack.push(val)}
        else this.minStack.push(Math.min(val, this.minStack.at(-1)))
        this.stack.push(val)
    }

    /**
     * @return {void}
     */
    pop() {
        if(this.minStack.length > 0) {
            this.minStack.pop()
        }
        if(this.stack.length > 0) this.stack.pop()
    }

    /**
     * @return {number}
     */
    top() {
       return this.stack.at(-1)
    }

    /**
     * @return {number}
     */
    getMin() {
        return this.minStack.at(-1)
    }
}
