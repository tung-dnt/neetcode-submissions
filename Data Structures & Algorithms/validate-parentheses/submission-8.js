class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isValid(s) {
        const dictionary = {
            "{": "}",
            "[": "]",
            "(": ")",
        };

        const stack = [];
        
        for(const char of s) {
            if(char in dictionary) {
                stack.push(char)
            }else {
                if(stack.length === 0) return false
                const lastOpen = stack.pop()
                if(dictionary[lastOpen] !== char) return false
            }
        }

        return stack.length === 0
    }
}
