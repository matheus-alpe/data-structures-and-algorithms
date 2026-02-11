package org.example.algorithms.sliding_window;

import java.util.HashSet;
import java.util.Set;

public class ContainsNearbyDuplicate {
    // https://leetcode.com/problems/contains-duplicate-ii/description
    public static boolean containsNearbyDuplicate(int[] nums, int k) {
        int left = 0;
        int right = 0;
        int max = nums.length - 1;

        while (right < max) {
            right++;
            int diff = Math.abs(left - right);
            if (nums[left] == nums[right] && diff <= k) return true;

            if (diff >= k || right == max) {
                right = ++left;
            }
        }

        return false;
    }

    public static boolean containsNearbyDuplicatePerformance(int[] nums, int k) {
        Set<Integer> set = new HashSet<Integer>(k);
        int i = 0, j = 0;
        int n = nums.length;
        while (j <= k && j < n) {
            if (!set.add(nums[j++])) {
                return true;
            }
        }

        while (j < n) {
            set.remove(nums[i++]);
            if (!set.add(nums[j++])) {
                return true;
            }
        }
        return false;
    }
}
