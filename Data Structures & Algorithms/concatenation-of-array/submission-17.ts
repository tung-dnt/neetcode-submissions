class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums: number[]): number[] {
        const ans: number[] = []

        for (let i = 0; i < nums.length * 2; i++) {
            ans.push(nums[i % nums.length])
        }

        return ans
    }
}
