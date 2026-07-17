class Solution {
    productExceptSelf(nums) {
        let allProduct = 1
        let zeroCount = 0   // count zeros, not just has0
        const result = []
        for (const num of nums) {
            if (num === 0) {
                zeroCount++
                continue
            }
            allProduct *= num
        }
        for (const num of nums) {
            if (num === 0) {
                result.push(zeroCount >= 2 ? 0 : allProduct)
                continue
            }
            result.push(zeroCount >= 1 ? 0 : allProduct / num)
        }
        return result
    }
}
