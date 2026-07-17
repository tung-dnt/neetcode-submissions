class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums) {
        const ans = [...nums]

        for (const num of nums) {
            ans.push(num)
        }

        return ans
    }
}
