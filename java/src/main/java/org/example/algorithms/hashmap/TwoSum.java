package org.example.algorithms.hashmap;

import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    // https://leetcode.com/problems/two-sum/description/
    public static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> hasher = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            if (hasher.containsKey(nums[i])) {
                return new int[]{hasher.get(nums[i]), i};
            }

            hasher.put(target - nums[i], i);
        }

        return new int[0];
    }
}
