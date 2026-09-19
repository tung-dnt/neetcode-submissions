class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s: string): boolean {
        let l = 0;
        let r = s.length - 1;

        while (l < r) {
            if (s[l].toLowerCase() === s[l].toUpperCase() && Number.isNaN(parseInt(s[l]))) {
                l++;
            } else if (s[r].toLowerCase() === s[r].toUpperCase() && Number.isNaN(parseInt(s[r]))) {
                r--;
            } else if (s[l].toLowerCase() !== s[r].toLowerCase()) {
                return false;
            } else {
                l++;
                r--;
            }
        }

        return true;
    }
}
