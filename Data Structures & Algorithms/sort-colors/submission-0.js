class Solution {
    /**
     * @param {number[]} nums
     * @return {void} Do not return anything, modify nums in-place instead.
     */
    sortColors(nums) {
        const buckets = []

        for (const num of nums) {
            buckets[num] ??= 0
            buckets[num] = buckets[num] + 1
        }


        let idx = 0
        for (let i = 0; i < buckets.length; i++) {
            for (let j = 0; j < buckets[i]; j++) {
                nums[idx] = i
                idx++
            }
        }

    }
}
