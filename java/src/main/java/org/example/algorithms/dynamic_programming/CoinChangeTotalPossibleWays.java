package org.example.algorithms.dynamic_programming;

public class CoinChangeTotalPossibleWays {
    // https://leetcode.com/problems/coin-change-ii/submissions/1902492332/
    public int change(int amount, int[] coins) {
        int[] dp = new int[amount + 1];
        dp[0] = 1; // There's one way to make amount 0: use no coins

        for (int coin : coins) {
            for (int x = coin; x <= amount; x++) {
                dp[x] += dp[x - coin];
            }
        }

        return dp[amount];
    }
}

