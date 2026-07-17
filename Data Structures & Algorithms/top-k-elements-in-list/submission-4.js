class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const count = {}
        for (const num of nums) count[num] = (count[num] ?? 0) + 1

        const results = Array.from({ length: k })

        for (const num in count) {
            const freq = count[num];
            const numVal = Number(num);
            const pair = { num: numVal, count: freq };

            const emptyIndex = results.findIndex((p) => p === undefined);

            if (emptyIndex !== -1) {
                results[emptyIndex] = pair;
            } else {
                let minIndex = 0;
                for (let i = 1; i < results.length; i++) {
                    if (results[i].count < results[minIndex].count) {
                        minIndex = i;
                    }
                }
                if (freq > results[minIndex].count) {
                    results[minIndex] = pair;
                }
            }
        }

        return results.map(pair => pair.num)
    }
}
