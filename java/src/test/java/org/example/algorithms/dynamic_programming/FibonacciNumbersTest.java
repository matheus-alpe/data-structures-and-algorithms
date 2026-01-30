package org.example.algorithms.dynamic_programming;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class FibonacciNumbersTest {
    @Test
    void testFibonacciRecursive() {
        FibonacciNumbers fibonacciNumbers = new FibonacciNumbers();
        assertEquals(55, fibonacciNumbers.fibonacci(10));
        assertEquals(1134903170, fibonacciNumbers.fibonacci(45));
    }

}