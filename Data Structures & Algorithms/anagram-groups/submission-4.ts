class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = new Map<string, string[]>()

        for (let str of strs) {
            const sorted = str.split('').sort().join('')

            if (map.has(sorted)) map.get(sorted).push(str)
            else map.set(sorted, [str])
        }

        return  Array.from(map.values());
    }
}
