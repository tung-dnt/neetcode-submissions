class Solution {
    /**
     * @param {number} x
     * @return {number}
     */
    mySqrt(x: number): number {
        if (x === 0) return 0;

        let [low, high] = [1, x]

        while (low <= high) {
            const mid = Math.floor((low + high) / 2)
            const exp = mid * mid

            if (exp < x) {
                low = mid + 1;
            } else if (exp > x) {
                high = mid - 1;
            } else {
                return mid;
            }
        }

        return high;
    }
}
