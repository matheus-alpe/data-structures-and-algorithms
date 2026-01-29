package org.example.algorithms.dynamic_programming;

public class ClimbingStairs {
    // do not use this method for large n due to exponential time complexity
    public int climbStairsRecursive(int n) {
        if (n <= 2) return n;
        return climbStairsRecursive(n - 1) + climbStairsRecursive(n - 2);
    }

    // https://leetcode.com/problems/climbing-stairs/description/
    public int climbStairs(int n) {
        if (n <= 2) return n;
        int[] dp = new int[n + 1];
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 3;
        for (int i = 4; i <= n; i++) {
            dp[i] = dp[i - 1] + dp[i - 2];
        }
        return dp[n];
    }
}
