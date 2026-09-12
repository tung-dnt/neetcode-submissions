class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s: string): boolean {
        const map = {
            "{": "}",
            "[": "]",
            "(": ")",
        };

        const stack: string[] = [];

        for (let i = 0; i < s.length; i++) {
            if (map[s[i]]) stack.push(s[i]);
            else {
                if (s[i] !== map[stack.pop()]) return false;
            }
        }

        return stack.length === 0;
    }
}
