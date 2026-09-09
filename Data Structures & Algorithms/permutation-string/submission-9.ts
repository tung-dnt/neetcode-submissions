class Solution {
    /**
     * @param {string} s1
     * @param {string} s2
     * @return {boolean}
     */
    checkInclusion(s1: string, s2: string): boolean {
        if (s1.length > s2.length) return false

        const s1FreqArr = new Array(26).fill(0)
        const s2FreqArr = new Array(26).fill(0)

        const isMatching = (): boolean => {
            for (let i = 0; i < 26; i++) {
                if (s1FreqArr[i] !== s2FreqArr[i]) return false
            }

            return true
        }

        for (let i = 0; i < s1.length; i++) {
            s1FreqArr[s1[i].charCodeAt(0) - 97] += 1
            s2FreqArr[s2[i].charCodeAt(0) - 97] += 1
        }

        if (isMatching()) return true

        for (let i = s1.length; i <= s2.length - 1; i++) {
            s2FreqArr[s2[i].charCodeAt(0) - 97] += 1
            s2FreqArr[s2[i - s1.length].charCodeAt(0) - 97] -= 1

            if (isMatching()) return true
        }

        return false
    }
}
