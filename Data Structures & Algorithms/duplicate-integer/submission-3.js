class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        const m = {}
        for (const num of nums) {
            const value = m[num] ?? 0
            if (value > 0) return true
            m[num] = value + 1
        }
        return false
    }
}
