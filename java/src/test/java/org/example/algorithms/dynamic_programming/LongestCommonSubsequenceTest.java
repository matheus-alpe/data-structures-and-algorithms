package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LongestCommonSubsequenceTest {
    @Test
    public void testLongestCommonSubsequence() {
        LongestCommonSubsequence lcs = new LongestCommonSubsequence();
        assertEquals(3, lcs.longestCommonSubsequence("abcde", "ace"));
        assertEquals(3, lcs.longestCommonSubsequence("abc", "abc"));
        assertEquals(0, lcs.longestCommonSubsequence("abc", "def"));
        assertEquals(4, lcs.longestCommonSubsequence("AGGTAB", "GXTXAYB"));
        assertEquals(1, lcs.longestCommonSubsequence("ab", "ba"));
    }
}