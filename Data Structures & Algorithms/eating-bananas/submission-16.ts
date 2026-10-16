class Solution {
    /**
     * @param {number[]} piles
     * @param {number} h
     * @return {number}
     */
    minEatingSpeed(piles: number[], h: number): number {
        let max = 0;
        for (let pile of piles) {
            if (pile > max) max = pile
        }

        let [l, r] = [1, max]

        const calculateTime = (piles: number[], k: number, h: number) => {
            let sum = 0;
            for (let pile of piles) {
                sum += Math.ceil(pile / k)

                if (sum > h) return sum;
            }

            return sum;
        }

        while (l <= r) {
            const mid = Math.floor((l + r) / 2);
            const time = calculateTime(piles, mid, h);

            if (time > h) {
                l = mid + 1;
            } else if (time <= h) {
                r = mid - 1;
            }
        }

        return l;
    }
}
