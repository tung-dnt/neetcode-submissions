/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * function guess(num) {}
 */

class Solution {
    /**
     * @param {number} n
     * @return {number}
     */
    guessNumber(n: number): number {
        let [low, high] = [1, n];

        while (low <= high) {
            const mid = Math.floor((low + high) / 2)

            switch (guess(mid)) {
                case 1:
                    low = mid + 1;
                    break;
                case -1:
                    high = mid - 1;
                    break;
                default:
                    return mid;
            }
        }
    }
}
