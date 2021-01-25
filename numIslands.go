package main

import "fmt"

/**
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。
示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands

 */

func numIslands(grid [][]byte) int {
	if len(grid) == 0{
		return 0
	}
	row, col := len(grid), len(grid[0])
	var count int

	for x := 0; x < row; x++{
		for y := 0; y < col; y++{
			if grid[x][y] == '1'{
				count++
				dfs(x, y, grid)
			}
		}
	}
	return count
}

func dfs(x, y int, grid [][]byte){
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0'{
		return
	}
	grid[x][y] = '0'
	dfs(x - 1, y, grid)
	dfs(x + 1, y, grid)
	dfs(x, y - 1, grid)
	dfs(x, y + 1, grid)
}

func main()  {
	test := [][]byte{
	{'1','1','1','1','0'},
	{'1','0','0','1','0'},
	{'0','1','0','0','0'},
	{'0','0','0','1','1'},
	}
	res := numIslands(test)
	fmt.Print(res)
}