package org.example.algorithms.sliding_window;

public class LongestSubstringWithoutRepeatingCharacters {
    // https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
    public static int lengthOfLongestSubstring(String s) {
        int left = 0;
        int max = 0;
        int[] frequency = new int[128];

        for (int right = 0; right < s.length(); right++) {
            frequency[s.charAt(right)]++;

            while(frequency[s.charAt(right)] > 1) {
                frequency[s.charAt(left)]--;
                left++;
            }

            max = Math.max(max, right - left + 1);
        }

        return max;
    }
}
