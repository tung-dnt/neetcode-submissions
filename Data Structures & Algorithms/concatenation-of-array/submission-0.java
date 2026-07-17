class Solution {
    public int[] getConcatenation(int[] nums) {
        int[] ans = new int[nums.length * 2];
        for(int i = 0; i < ans.length; ++i) {
            int originalLength = nums.length;
            int originalIdx = i < originalLength ? i : i - originalLength;
            ans[i] = nums[originalIdx];
        }
        return ans;
    }
}