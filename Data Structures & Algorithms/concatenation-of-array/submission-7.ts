class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums: number[]): number[] {
        const ans = new Array(nums.length * 2)

        nums.forEach((n, i) => {
            ans[i] = n
            ans[i + nums.length] = n
        })

        return ans
    }
}
