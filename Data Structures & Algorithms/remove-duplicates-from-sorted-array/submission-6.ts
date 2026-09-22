class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    removeDuplicates(nums: number[]): number {
        if (nums.length === 1) return 1

        let l = 0
        let r = 0

        while (r < nums.length) {
            nums[l] = nums[r]
            while (nums[l] === nums[r]) r++

            l++
        }

        return l
    }
}