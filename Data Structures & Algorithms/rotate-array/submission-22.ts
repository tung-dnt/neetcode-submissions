class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {void} Do not return anything, modify nums in-place instead.
     */
    rotate(nums: number[], k: number): void {
        const n = nums.length
        k %= n

        const reverse = (l: number, r: number) => {
            while (l < r) {
                const temp = nums[l]
                nums[l] = nums[r]
                nums[r] = temp

                l++;
                r--;
            }
        }

        reverse(0, n - 1)
        reverse(0, k - 1)
        reverse(k, n - 1)
    }
}
