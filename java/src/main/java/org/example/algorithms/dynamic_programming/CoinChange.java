package org.example.algorithms.dynamic_programming;

public class CoinChange {
    // https://leetcode.com/problems/coin-change/description/
    public int coinChange(int[] coins, int amount) {
        int[] dp = new int[amount + 1];
        // Initialize dp array with a value greater than any possible number of coins
        for (int i = 1; i <= amount; i++) {
            dp[i] = Integer.MAX_VALUE;
        }
        for (int coin : coins) {
            for (int x = coin; x <= amount; x++) {
                if (x == coin) {
                    // Base case: if the amount is exactly equal to the coin value
                    dp[x] = Math.min(dp[x], 1);
                } else if (dp[x - coin] != Integer.MAX_VALUE) {
                    // If we can make change for (x - coin), update dp[x]
                    dp[x] = Math.min(dp[x], dp[x - coin] + 1);
                }
            }
        }
        return dp[amount] == Integer.MAX_VALUE ? -1 : dp[amount];
    }
}
