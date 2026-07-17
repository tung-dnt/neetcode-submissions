class Solution {
    productExceptSelf(nums) {
        // Division

        // let allProduct = 1
        // let zeroCount = 0   // count zeros, not just has0
        // const result = []
        // for (const num of nums) {
        //     if (num === 0) {
        //         zeroCount++
        //         continue
        //     }
        //     allProduct *= num
        // }
        // for (const num of nums) {
        //     if (num === 0) {
        //         result.push(zeroCount >= 2 ? 0 : allProduct)
        //         continue
        //     }
        //     result.push(zeroCount >= 1 ? 0 : allProduct / num)
        // }
        // return result

        // Prefix & Suffix
        const result = Array(nums.length)
        result[0] = 1

        // Prefix calculation
        for (let i = 1; i < nums.length; i++) {
            result[i] = result[i - 1] * nums[i - 1]
        }

        let suffix = nums[nums.length - 1]
        for (let i = nums.length - 2; i >= 0; i--) {
            result[i] *= suffix
            suffix *= nums[i]
        }

        return result
    }
}
