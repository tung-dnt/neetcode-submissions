class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    removeDuplicates(nums: number[]): number {
        const set = new Set(nums);
        let i = 0;
        set.forEach((n) => {
            nums[i++] = n;
        });

        return set.size;
    }
}