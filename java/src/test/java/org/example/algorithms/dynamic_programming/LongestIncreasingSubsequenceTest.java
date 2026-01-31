package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class LongestIncreasingSubsequenceTest {
    @Test
    public void testLengthOfLIS() {
        LongestIncreasingSubsequence lis = new LongestIncreasingSubsequence();

        assertEquals(4, lis.lengthOfLIS(new int[]{10, 9, 2, 5, 3, 7, 101, 18}));
        assertEquals(4, lis.lengthOfLIS(new int[]{0, 1, 0, 3, 2, 3}));
        assertEquals(1, lis.lengthOfLIS(new int[]{7, 7, 7, 7, 7, 7, 7}));
        assertEquals(5, lis.lengthOfLIS(new int[]{1, 2, 3, 4, 5}));
        assertEquals(1, lis.lengthOfLIS(new int[]{5, 4, 3, 2, 1}));
    }
}