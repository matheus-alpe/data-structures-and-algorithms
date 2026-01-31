package org.example.algorithms.dynamic_programming;

import java.util.Arrays;

public class LastStoneWeight2 {
    // https://leetcode.com/problems/last-stone-weight-ii/description/
    public int lastStoneWeightII(int[] stones) {
        int totalWeight = 0;
        for (int stone : stones) {
            totalWeight += stone;
        }
        int target = totalWeight / 2;

        boolean[] dp = new boolean[target + 1];
        dp[0] = true;

        for (int stone : stones) {
            for (int j = target; j >= stone; j--) {
                dp[j] = dp[j] || dp[j - stone];
            }
        }

        for (int i = target; i >= 0; i--) {
            if (dp[i]) {
                return totalWeight - 2 * i;
            }
        }
        return 0;
    }
}
