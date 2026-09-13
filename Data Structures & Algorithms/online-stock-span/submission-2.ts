class StockSpanner {
    constructor() {}

    readonly prices: number[] = [];

    /**
     * @param {number} price
     * @return {number}
     */
    next(price: number): number {
        let count: number = 1;

        for (let i = this.prices.length - 1; i >= 0; i--) {
            if (this.prices[i] <= price) count++;
            else break;
        }

        this.prices.push(price);

        return count;
    }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * var obj = new StockSpanner()
 * var param_1 = obj.next(price)
 */
