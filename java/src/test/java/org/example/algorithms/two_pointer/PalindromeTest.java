package org.example.algorithms.two_pointer;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class PalindromeTest {
    @Test
    public void testIsPalindrome() {
        Palindrome palindrome = new Palindrome();

        assertTrue(palindrome.isPalindrome("A man, a plan, a canal: Panama"));
        assertFalse(palindrome.isPalindrome("race a car"));
        assertTrue(palindrome.isPalindrome(" "));
        assertTrue(palindrome.isPalindrome("a."));
        assertFalse(palindrome.isPalindrome("0P"));
    }
}