class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s: string): boolean {
        const l: number = s.length;

        const map = new Map();
        map.set("{", "}");
        map.set("[", "]");
        map.set("(", ")");

        if (l % 2 !== 0) return false;

        const stack: string[] = [];

        for (let i = 0; i < l; i++) {
            if (s[i] === "{" || s[i] === "[" || s[i] === "(") stack.push(s[i]);
            else {
                if (s[i] !== map.get(stack.pop())) return false;
            }
        }

        if (stack.length > 0) return false;

        return true;
    }
}