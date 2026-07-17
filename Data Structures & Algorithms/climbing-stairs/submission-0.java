class Solution {
    private int climb(int n, int step) {
        if(step >= n) return step == n ? 1 : 0;
        return climb(n, step + 1) + climb(n, step + 2);
    }
    public int climbStairs(int n) {
        return climb(n, 0);
    }
}
