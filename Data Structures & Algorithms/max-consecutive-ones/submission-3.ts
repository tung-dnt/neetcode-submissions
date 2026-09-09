class Solution {
  /**
   * @param {number[]} nums
   * @return {number}
   */
  findMaxConsecutiveOnes(nums: number[]): number {
    let count = 0;
    let max = 0;

    nums.forEach((n) => {
      if (n === 1) {
        count += 1;
      } else {
        if (count > max) {
          max = count;
        }

        count = 0;
      }
    });

    if (count > max) max = count

    return max;
  }
}
