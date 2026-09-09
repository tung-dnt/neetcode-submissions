class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findMaxConsecutiveOnes(nums) {
        let result = 0;
        let max = 0;
        nums.forEach((num, i) => {
            if (num == 1) result++;
            else {
                result = 0;
            }
                max = Math.max(result, max);
        });
        return max;
    }
}
