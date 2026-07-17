class Solution {
    /**
     * @param {number[]} piles
     * @param {number} h
     * @return {number}
     */
    minEatingSpeed(piles, h) {
        let speed = 1
        while (true) {
            let totalTimes = 0
            for (let pile of piles) {
                totalTimes += Math.ceil(pile / speed)
            }

            if (totalTimes <= h) {
                return speed
            }

            speed++
        }
    }
}
