class Solution {
    /**
     * @param {number[][]} matrix
     * @param {number} target
     * @return {boolean}
     */
    searchMatrix(matrix: number[][], target: number): boolean {
        if (target < matrix[0][0]) return false;

        let [top, down, left, right] = [0, matrix.length - 1, 0, matrix[0].length - 1];
        let rowIdx;

        while (top <= down) {
            const mid = Math.floor((top + down) / 2);

            if (matrix[mid][0] > target) {
                down = mid - 1;
            } else if (matrix[mid][0] < target) {
                top = mid + 1;
            } else {
                return true;
            }
        }

        rowIdx = down;

        while (left <= right) {
            const mid = Math.floor((left + right) / 2);

            if (matrix[rowIdx][mid] > target) {
                right = mid - 1;
            } else if (matrix[rowIdx][mid] < target) {
                left = mid + 1;
            } else {
                return true;
            }
        }

        return false;
    }
}
