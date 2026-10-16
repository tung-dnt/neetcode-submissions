class Solution {
    /**
     * @param {number[]} piles
     * @param {number} h
     * @return {number}
     */
    minEatingSpeed(piles: number[], h: number): number {
        const max = Math.max(...piles)
        let [l, r] = [1, max]

        const calculateTime = (piles: number[], k: number) => {
            let sum = 0;
            for (let pile of piles) {
                sum += Math.ceil(pile / k)
            }

            return sum;
        }

        let minK: number;

        while (l <= r) {
            const mid = Math.floor((l + r) / 2);
            const time = calculateTime(piles, mid);

            if (time > h) {
                l = mid + 1;
            } else if (time <= h) {
                r = mid - 1;

                if (!minK || minK > mid) minK = mid
            }
        }

        return minK;
    }
}
