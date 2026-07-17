class Solution {
    /**
     * For each number, find the next number 
     * 
     * @param {number[]} nums
     * @return {number}
     */
    longestConsecutive(nums) {
        const numSet = new Set(nums)
        let longest = 0

        for (const num of numSet) {
            // Make sure current number is the smallest value of consecutive sequence
            // by making sure there's no smaller item
            if (numSet.has(num - 1)) continue
            let length = 1
            while (numSet.has(num + length)) length++
            longest = Math.max(length, longest)
        }

        return longest
    }
}
