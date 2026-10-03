class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = new Map<string, string[]>();

        for (let str of strs) {
            const freqs = new Array(26).fill(0);

            for (let char of str) {
                freqs[char.charCodeAt(0) - "a".charCodeAt(0)]++;
            }

            const freqsString = freqs.join(",");

            if (map.has(freqsString)) map.get(freqsString).push(str);
            else map.set(freqsString, [str]);
        }

        return Array.from(map.values());
    }
}
