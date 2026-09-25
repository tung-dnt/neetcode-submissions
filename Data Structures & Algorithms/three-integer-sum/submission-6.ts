class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    threeSum(nums: number[]): number[][] {
        nums.sort((a, b) => a - b);

        const results: number[][] = [];

        nums.forEach((num, i) => {
            if (i > 0 && nums[i] === nums[i - 1]) return

            let l = i + 1;
            let r = nums.length - 1;

             while (l < r) {
                if (nums[l] + nums[r] + num === 0) {
                    results.push([num, nums[l], nums[r]]);
                    l++;
                    r--;

                    while (l < r && nums[l] === nums[l - 1]) l++;
                } else if (nums[l] + nums[r] + num < 0) {
                    l++;
                } else {
                    r--;
                }
            }
        });

        return results;
    }
}
