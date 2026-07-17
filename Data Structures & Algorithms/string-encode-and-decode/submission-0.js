class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(strs) {
        if (strs.length === 0) return ''

        let result = ''
        for (const str of strs) {
            result += str.length + '#' + str
        }
        return result
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        const result = []
        let startIdx = 0
        while (startIdx < str.length) {
            let endIdx = startIdx
            // Find length index of orginal substring
            while (str[endIdx] !== '#') endIdx++

            // Once # delimeter found, get length of substring characters
            const length = parseInt(str.substring(startIdx, endIdx))
            // Skip # character and increase starting index by 1
            startIdx = endIdx + 1
            endIdx = startIdx + length

            result.push(str.substring(startIdx, endIdx))

            startIdx = endIdx
        }
        return result
    }
}