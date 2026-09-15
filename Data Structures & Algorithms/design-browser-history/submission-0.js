class BrowserHistory {
    /**
     * @constructor
     * @param {string} homepage
     */
    constructor(homepage) {
        this.stack = [homepage];
        this.curr = 0;
    }

    /**
     * @param {string} url
     * @return {void}
     */
    visit(url) {
        this.stack = this.stack.slice(0, this.curr + 1);
        this.stack.push(url);
        this.curr = this.stack.length - 1;
    }

    /**
     * @param {number} steps
     * @return {string}
     */
    back(steps) {
        this.curr = Math.max(0, this.curr - steps);
        return this.stack[this.curr];
    }

    /**
     * @param {number} steps
     * @return {string}
     */
    forward(steps) {
        this.curr = Math.min(this.stack.length - 1, this.curr + steps);
        return this.stack[this.curr];
    }
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * var obj = new BrowserHistory(homepage)
 * obj.visit(url)
 * var param_2 = obj.back(steps)
 * var param_3 = obj.forward(steps)
 */
