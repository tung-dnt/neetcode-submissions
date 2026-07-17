class Solution {
    public int[] getConcatenation(int[] nums) {
        int originalLength = nums.length;
        int[] ans = new int[originalLength * 2];
        for(int i = 0; i < ans.length; ++i) {
            int originalIdx = i < originalLength ? i : i - originalLength;
            ans[i] = nums[originalIdx];
        }
        return ans;
    }
}