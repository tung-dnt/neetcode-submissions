class Solution {
    /**
     * @param {number[]} heights
     * @return {number}
     */
    maxArea(heights: number[]): number {
        let l = 0;
        let r = heights.length - 1;

        let max = 0;

        while (l < r) {
            const minHeight = Math.min(heights[l], heights[r]);

            const area = (r - l) * minHeight;

            if (area > max) max = area;

            if (heights[l] === minHeight) l++;
            else r--;
        }

        return max;
    }
}
