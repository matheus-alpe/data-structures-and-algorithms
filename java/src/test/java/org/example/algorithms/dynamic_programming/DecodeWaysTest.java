package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class DecodeWaysTest {
    @Test
    public void testNumDecodings() {
        DecodeWays decodeWays = new DecodeWays();

        assertEquals(2, decodeWays.numDecodings("12")); // "AB" (1 2) or "L" (12)
        assertEquals(3, decodeWays.numDecodings("226")); // "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6)
        assertEquals(0, decodeWays.numDecodings("0")); // No valid decodings
        assertEquals(0, decodeWays.numDecodings("06")); // No valid decodings
        assertEquals(1, decodeWays.numDecodings(""));
        assertEquals(1, decodeWays.numDecodings("10")); // "J" (10)
        assertEquals(1, decodeWays.numDecodings("27")); // "BG" (2 7)
        assertEquals(3, decodeWays.numDecodings("1234")); // "ABCD", "LCD", "AWD"
    }
}