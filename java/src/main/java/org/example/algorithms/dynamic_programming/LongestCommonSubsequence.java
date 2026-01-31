package org.example.algorithms.dynamic_programming;

public class LongestCommonSubsequence {
    // https://leetcode.com/problems/longest-common-subsequence/description/
    public int longestCommonSubsequenceMySolution(String text1, String text2) {
        int m = text1.length();
        int n = text2.length();
        int[][] dp = new int[m + 1][n + 1];

        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (text1.charAt(i - 1) == text2.charAt(j - 1)) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                }
            }
        }
        return dp[m][n];
    }

    public int longestCommonSubsequence(String text1, String text2) {
        int[] dp = new int[text1.length()];
        int longest = 0;

        for (char c : text2.toCharArray()) {
            int currentLength = 0;
            for (int i = 0; i < dp.length; i++) {
                if (currentLength < dp[i]) {
                    currentLength = dp[i];
                } else if (c == text1.charAt(i)) {
                    dp[i] = currentLength + 1;
                    longest = Math.max(longest, dp[i]);
                }
            }
        }

        return longest;
    }

}
