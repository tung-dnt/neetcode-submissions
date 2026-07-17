class Solution {
    public int removeElement(int[] nums, int val) {
        int unmatchedCursor = 0;
        for(int i = 0; i < nums.length; ++i) {
            if(nums[i] != val) {
                nums[unmatchedCursor] = nums[i];
                unmatchedCursor++;
            }
        }

        return unmatchedCursor;
    }
}