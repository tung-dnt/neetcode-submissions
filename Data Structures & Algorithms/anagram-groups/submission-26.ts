class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs: string[]): string[][] {
        const map = {};

        for (let str of strs) {
            const freqs = new Array(26).fill(0);

            for (let char of str) {
                freqs[char.charCodeAt(0) - "a".charCodeAt(0)]++;
            }

            const freqsString = freqs.join(",");

            if (map[freqsString]) map[freqsString].push(str);
            else map[freqsString] = [str];
        }

        return Object.values(map);
    }
}
