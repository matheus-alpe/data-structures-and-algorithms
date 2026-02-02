package org.example.algorithms.backtracking;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class LetterCombinationsPhoneNumberTest {
    @Test
    public void letterCombinationsTest() {
        LetterCombinationsPhoneNumber solver = new LetterCombinationsPhoneNumber();

        assertArrayEquals(
                new String[]{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
                solver.letterCombinations("23").toArray()
        );

        assertArrayEquals(
                new String[]{"a", "b", "c"},
                solver.letterCombinations("2").toArray()
        );

        assertArrayEquals(
                new String[]{},
                solver.letterCombinations("").toArray()
        );
    }

}