package graph

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == '1' {
				islands++
				dfs(grid, x, y)
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, x int, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}

	// mark as visted
	grid[x][y] = '0'

	// explore 4 directions
	dfs(grid, x+1, y) // down
	dfs(grid, x-1, y) // update
	dfs(grid, x, y+1) // right
	dfs(grid, x, y-1) // left
}
