package org.example.algorithms.dynamic_programming;

public class LongestIncreasingSubsequence {
    // https://leetcode.com/problems/longest-increasing-subsequence/description/
    public int lengthOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        int longest = 0;

        for (int i = 0; i < nums.length; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
            longest = Math.max(dp[i], longest);
        }

        return longest;
    }
}
