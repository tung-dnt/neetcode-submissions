class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        let queue = '', stack = '',tail = s.length - 1,head = 0

        while (tail > -1 && head < s.length) {
            if (head < s.length) {
                if (/\w/.test(s[head])) queue += s[head].toLowerCase()
                head++
            }
            if (tail > -1) {
                if(/\w/.test(s[tail])) stack += s[tail].toLowerCase()
                tail--
            }
        }

        return stack === queue
    }
}
