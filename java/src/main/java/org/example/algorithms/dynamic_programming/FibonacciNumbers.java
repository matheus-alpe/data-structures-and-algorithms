package org.example.algorithms.dynamic_programming;

public class FibonacciNumbers {
    public int fibonacci(int n) {
       int[]  dp = new int[n + 1];
       dp[0] = 0;
       dp[1] = 1;
       return fibonacciRecursive(n, dp);
    }

    private int fibonacciRecursive(int n, int[] dp) {
        if (dp[n] != 0 || n <= 1) return dp[n];
        dp[n] = fibonacciRecursive(n - 1, dp) + fibonacciRecursive(n - 2, dp);
        return dp[n];
    }
}
