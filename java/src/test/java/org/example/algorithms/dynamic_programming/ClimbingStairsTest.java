package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ClimbingStairsTest {
    @Test
    void testClimbStairs() {
        ClimbingStairs cs = new ClimbingStairs();
        assertEquals(1, cs.climbStairs(1));
        assertEquals(2, cs.climbStairs(2));
        assertEquals(3, cs.climbStairs(3));
        assertEquals(5, cs.climbStairs(4));
        assertEquals(8, cs.climbStairs(5));
        assertEquals(13, cs.climbStairs(6));
        assertEquals(21, cs.climbStairs(7));
        assertEquals(34, cs.climbStairs(8));
    }

    @Test
    void testClimbStairsEdgeCases() {
        ClimbingStairs cs = new ClimbingStairs();
        assertEquals(1836311903, cs.climbStairs(45));// edge case for large number of steps
    }
}