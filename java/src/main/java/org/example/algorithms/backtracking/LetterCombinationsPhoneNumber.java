package org.example.algorithms.backtracking;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class LetterCombinationsPhoneNumber {
    static HashMap<Character, char[]> digitToLetters = new HashMap<>() {{
        put('2', new char[]{'a', 'b', 'c'});
        put('3', new char[]{'d', 'e', 'f'});
        put('4', new char[]{'g', 'h', 'i'});
        put('5', new char[]{'j', 'k', 'l'});
        put('6', new char[]{'m', 'n', 'o'});
        put('7', new char[]{'p', 'q', 'r', 's'});
        put('8', new char[]{'t', 'u', 'v'});
        put('9', new char[]{'w', 'x', 'y', 'z'});
    }};

    // https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
    public List<String> letterCombinations(String digits) {
        if (digits.isEmpty()) return List.of();

        return backtrack(
                digits,
                0,
                new StringBuilder(),
                new ArrayList<>()
        );
    }

    private List<String> backtrack(String digits, int index, StringBuilder current, List<String> result) {
        if (index == digits.length()) {
            result.add(current.toString());
            return result;
        }

        char[] letters = digitToLetters.get(digits.charAt(index) );
        for (char letter : letters) {
            current.append(letter);
            backtrack(digits, index + 1, current, result);
            current.deleteCharAt(current.length() - 1);
        }

        return result;
    }
}
