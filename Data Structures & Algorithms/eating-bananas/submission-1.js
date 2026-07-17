class Solution {
    /**
     * @param {number[]} piles
     * @param {number} h
     * @return {number}
     */
    minEatingSpeed(piles, h) {
        let left = 1;
        let right = Math.max(...piles) 

        let result = right;

        while (left <= right) {
            const mid = Math.floor((left + right) / 2)
            let totalTimes = 0
            for (const pile of piles) {
                totalTimes += Math.ceil(pile / mid)
            }

            if (totalTimes <= h) {
                result = mid
                right = mid - 1
            } else {
                left = mid + 1
            }

        }

        return result
    }
}
