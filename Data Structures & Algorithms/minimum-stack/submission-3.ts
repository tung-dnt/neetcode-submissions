class MinStack {
    readonly arr: number[] = []
    readonly minArr: number[] = []    

    constructor() {}

    /**
     * @param {number} val
     * @return {void}
     */
    push(val: number): void {
        this.arr.push(val)
        
        if (this.minArr.length === 0 || val <= this.minArr[this.minArr.length - 1]) this.minArr.push(val)
    }

    /**
     * @return {void}
     */
    pop(): void {
        const pop = this.arr.pop()

        if (pop === this.minArr[this.minArr.length - 1]) this.minArr.pop()
    }

    /**
     * @return {number}
     */
    top(): number {
        return this.arr[this.arr.length - 1]
    }

    /**
     * @return {number}
     */
    getMin(): number {
        return this.minArr[this.minArr.length - 1]
    }
}