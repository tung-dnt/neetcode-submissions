class Solution {
    /**
     * @param {number[][]} matrix
     * @param {number} target
     * @return {boolean}
     */
    searchMatrix(matrix, target) {
        const colNums = matrix[0].length - 1
        let firstRow = 0
        let lastRow = matrix.length - 1


        let correctRow = 0

        while (firstRow <= lastRow) {

            const midRow = Math.floor((firstRow + lastRow) / 2)

            const firstVal = matrix[midRow][0]
            const lastVal = matrix[midRow][colNums]

            if (target === firstVal || target === lastVal) {
                return true
            }

            if (target > firstVal && target < lastVal) {
                correctRow = midRow
                break
            }

            if (target < firstVal) {
                lastRow = midRow - 1
            }

            if (target > lastVal) {
                firstRow = midRow + 1
            }
        }


        let left = 0
        let right = colNums

        while (left <= right) {
            const mid = Math.floor((left + right) / 2)

            const val = matrix[correctRow][mid]

            if (val === target) {
                return true
            } else if (target < val) {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }


        return false
    }
}
