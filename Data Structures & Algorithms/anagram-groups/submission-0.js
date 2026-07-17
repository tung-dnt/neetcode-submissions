class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     * 
     * -> Sort string to build map key -> group similar key
     */
    groupAnagrams(strs) {
        const res = {}
        for (const s of strs) {
            const sorted = s.split('').sort().join('')
            if (!res[sorted]) {
                res[sorted] = [];
            }
            res[sorted].push(s)
        }

        return Object.values(res)
    }
}
