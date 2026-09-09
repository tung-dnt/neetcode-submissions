class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums: number[]): number[] {
        const ans: number[] = nums

        nums.forEach((n) => ans.push(n))
        
        return ans
    }
}
