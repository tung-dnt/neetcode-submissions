class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    removeDuplicates(nums) {
        let nextNonDupe = 1;
        for(let i = 1; i < nums.length; ++i) {
            if(nums[nextNonDupe - 1] !== nums[i]) {
                nums[nextNonDupe] = nums[i]
                nextNonDupe++
            }
        }

        return nextNonDupe
    }
}
