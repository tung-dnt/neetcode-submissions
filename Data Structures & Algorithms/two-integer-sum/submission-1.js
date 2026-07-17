class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let left = 0
        while (left < nums.length - 1) {
            let right = left + 1
            while(right < nums.length) {
                if (nums[right] + nums[left] === target) 
                    return [left, right]
                right++    
            }
            left++
        }
        return []
    }
}
