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

        if (s.length % 2 !== 0) return false;

        const stack: string[] = [];

        for (let i = 0; i < s.length; i++) {
            if (s[i] === "{" || s[i] === "[" || s[i] === "(") stack.push(s[i]);
            else {
                if (s[i] !== map[stack.pop()]) return false;
            }
        }

        return stack.length === 0;
    }
}
