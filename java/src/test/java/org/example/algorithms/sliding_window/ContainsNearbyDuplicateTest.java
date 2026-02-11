package org.example.algorithms.sliding_window;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class ContainsNearbyDuplicateTest {
    @Test
    void testCase1() {
        int[] nums = {1, 2, 3, 1};
        int k = 3;
        assertEquals(true, ContainsNearbyDuplicate.containsNearbyDuplicate(nums, k));
    }

    @Test
    void testCase2() {
        int[] nums = {1, 0, 1, 1};
        int k = 1;
        assertEquals(true, ContainsNearbyDuplicate.containsNearbyDuplicate(nums, k));
    }

    @Test
    void testCase3() {
        int[] nums = {1,2,3,1,2,3};
        int k = 2;
        assertEquals(false, ContainsNearbyDuplicate.containsNearbyDuplicate(nums, k));
    }

    @Test
    void testCase4() {
        int[] nums = {1,2,3,4,5,6,7,8,9,9};
        int k = 3;
        assertEquals(true, ContainsNearbyDuplicate.containsNearbyDuplicate(nums, k));
    }

    @Test
    void testCase4Performance() {
        int[] nums = {1,2,3,4,5,6,7,8,9,9};
        int k = 3;
        assertEquals(true, ContainsNearbyDuplicate.containsNearbyDuplicatePerformance(nums, k));
    }

    @Test
    void testContainsDuplicateWithDuplicates() {
        org.example.algorithms.hashmap.ContainsDuplicate solution = new org.example.algorithms.hashmap.ContainsDuplicate();
        int[] nums = {1, 2, 3, 1};
        assertTrue(solution.containsDuplicate(nums));
    }

    @Test
    void testContainsDuplicateNoDuplicates() {
        org.example.algorithms.hashmap.ContainsDuplicate solution = new org.example.algorithms.hashmap.ContainsDuplicate();
        int[] nums = {1, 2, 3, 4};
        assertFalse(solution.containsDuplicate(nums));
    }

    @Test
    void testContainsDuplicateEmptyArray() {
        org.example.algorithms.hashmap.ContainsDuplicate solution = new org.example.algorithms.hashmap.ContainsDuplicate();
        int[] nums = {};
        assertFalse(solution.containsDuplicate(nums));
    }

    @Test
    void testContainsDuplicateSingleElement() {
        org.example.algorithms.hashmap.ContainsDuplicate solution = new org.example.algorithms.hashmap.ContainsDuplicate();
        int[] nums = {1};
        assertFalse(solution.containsDuplicate(nums));
    }

    @Test
    void testContainsDuplicateAllSame() {
        org.example.algorithms.hashmap.ContainsDuplicate solution = new org.example.algorithms.hashmap.ContainsDuplicate();
        int[] nums = {5, 5, 5, 5};
        assertTrue(solution.containsDuplicate(nums));
    }
}