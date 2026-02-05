package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class HouseRobberTest {
    @Test
    public void robTest() {
        HouseRobber solver = new HouseRobber();

        assertEquals(4, solver.rob(new int[]{1, 2, 3, 1}));
        assertEquals(12, solver.rob(new int[]{2, 7, 9, 3, 1}));
        assertEquals(0, solver.rob(new int[]{}));
        assertEquals(3, solver.rob(new int[]{3}));
        assertEquals(3, solver.rob(new int[]{2, 3}));
    }
}