class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const count = {}
        for (const num of nums) count[num] = (count[num] ?? 0) + 1
        const arr = Object.entries(count)
        arr.sort((a,b) => b[1] - a[1])

        return arr.slice(0, k).map(val => val[0])
    }
}
