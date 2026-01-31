package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class CoinChangeTotalPossibleWaysTest {
    @Test
    public void testChange() {
        CoinChangeTotalPossibleWays coinChangeTotalPossibleWays = new CoinChangeTotalPossibleWays();

        int[] coins1 = {1, 2, 5};
        int amount1 = 5;
        assertEquals(4, coinChangeTotalPossibleWays.change(amount1, coins1));

        int[] coins2 = {2};
        int amount2 = 3;
        assertEquals(0, coinChangeTotalPossibleWays.change(amount2, coins2));

        int[] coins3 = {10};
        int amount3 = 10;
        assertEquals(1, coinChangeTotalPossibleWays.change(amount3, coins3));

        int[] coins4 = {1, 2, 3};
        int amount4 = 4;
        assertEquals(4, coinChangeTotalPossibleWays.change(amount4, coins4));

        int[] coins5 = {};
        int amount5 = 0;
        assertEquals(1, coinChangeTotalPossibleWays.change(amount5, coins5));
    }
}