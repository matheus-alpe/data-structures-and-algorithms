package org.example.algorithms.graph;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Set;

public class NumberOfIslands {
    // https://leetcode.com/problems/number-of-islands/
    public int mySolutionNumIslands(char[][] grid) {
        int islands = 0;
        int rows = grid.length;
        int cols = grid[0].length;
        int[][] directions = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        Set<String> visited = new HashSet<>();


        for (int r = 0; r < rows; r++) {
            for (int c = 0; c < cols; c++) {
                if (grid[r][c] == '1' && !visited.contains(cellKey(r, c))) {
                    islands++;
                    bfs(grid, r, c, visited, directions, rows, cols);
                }
            }
        }

        return islands;
    }

    private void bfs(char[][] grid, int r, int c, Set<String> visited, int[][] directions, int rows, int cols) {
        Queue<int[]> queue = new LinkedList<>();
        visited.add(cellKey(r, c));
        queue.add(new int[]{r, c});

        while (!queue.isEmpty()) {
            int[] cell = queue.poll();
            int row = cell[0];
            int col = cell[1];

            for (int[] dir : directions) {
                int newRow = row + dir[0];
                int newCol = col + dir[1];
                boolean withinBounds = newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols;

                if (withinBounds && grid[newRow][newCol] == '1' && !visited.contains(cellKey(newRow, newCol))) {
                    visited.add(cellKey(newRow, newCol));
                    queue.add(new int[]{newRow, newCol});
                }
            }
        }
    }

    private String cellKey(int r, int c) {
        return String.format("%dx%d", r, c);
    }

    /**
     * Best solution
     */
    public int numIslands(char[][] grid) {

        if (grid == null || grid.length == 0) return 0;

        int islands = 0;

        for (int x = 0; x < grid.length; x++) {
            for (int y = 0; y < grid[0].length; y++) {

                if (grid[x][y] == '1') {
                    islands++;
                    dfs(grid, x, y);
                }
            }
        }

        return islands;
    }

    private void dfs(char[][] grid, int x, int y) {

        // Boundary or water check
        if (x < 0 || y < 0 ||
                x >= grid.length || y >= grid[0].length ||
                grid[x][y] == '0') {
            return;
        }

        // Mark visited
        grid[x][y] = '0';

        // Explore 4 directions
        dfs(grid, x + 1, y); // down
        dfs(grid, x - 1, y); // up
        dfs(grid, x, y + 1); // right
        dfs(grid, x, y - 1); // left
    }
}
