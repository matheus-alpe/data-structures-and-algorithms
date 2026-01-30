package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class CoinChangeTest {
    @Test
    public void testCoinChange() {
        CoinChange coinChange = new CoinChange();

        assertEquals(3, coinChange.coinChange(new int[]{1, 2, 5}, 11));
        assertEquals(-1, coinChange.coinChange(new int[]{2}, 3));
        assertEquals(0, coinChange.coinChange(new int[]{1}, 0));
        assertEquals(2, coinChange.coinChange(new int[]{1, 3, 4}, 6));
        assertEquals(2, coinChange.coinChange(new int[]{1, 5, 10, 25}, 30));
    }
}