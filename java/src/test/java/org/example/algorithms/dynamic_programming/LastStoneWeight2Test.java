package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class LastStoneWeight2Test {
    @Test
    public void testCase1() {
        LastStoneWeight2 solution = new LastStoneWeight2();
        assertEquals(1, solution.lastStoneWeightII(new int[]{2, 7, 4, 1, 8, 1}));
        assertEquals(5, solution.lastStoneWeightII(new int[]{31, 26, 33, 21, 40}));
    }
}