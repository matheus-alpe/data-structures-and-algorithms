package org.example.algorithms.sliding_window;

import java.util.HashMap;

public class LongestSubstringWithoutRepeatingCharacters {
    // https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
    public static int lengthOfLongestSubstring(String s) {
        int left = 0;
        int max = 0;
        HashMap<Character, Integer> frequency = new HashMap<>();

        for (int right = 0; right < s.length(); right++) {
            if (frequency.containsKey(s.charAt(right)) && frequency.get(s.charAt(right)) >= left) {
                left = frequency.get(s.charAt(right)) + 1;
            }

            max = Math.max(max, right - left + 1);
            frequency.put(s.charAt(right), right);
        }

        return max;
    }
}
