package main

import "fmt"

func main() {
	maze := [][]int{
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 0, 0, 0},
		{0, 1, 0, 0},
	}
	start := [2]int{0, 0}
	end := [2]int{3, 3}
	path := findPath(maze, start, end)
	if len(path) > 0 {
		fmt.Println("找到了一条路径：", path)
	} else {
		fmt.Println("没有找到路径。")
	}
}

func findPath(maze [][]int, start, end [2]int) []string {
	var dfs func(i, j int, path []string) []string
	dfs = func(i, j int, path []string) []string {
		if i == end[0] && j == end[1] {
			return path
		}
		if maze[i][j] == 1 {
			return nil
		}
		maze[i][j] = 1
		if i > 0 && maze[i-1][j] == 0 {
			if res := dfs(i-1, j, append(path, "上")); res != nil {
				return res
			}
		}
		if i < len(maze)-1 && maze[i+1][j] == 0 {
			if res := dfs(i+1, j, append(path, "下")); res != nil {
				return res
			}
		}
		if j > 0 && maze[i][j-1] == 0 {
			if res := dfs(i, j-1, append(path, "左")); res != nil {
				return res
			}
		}
		if j < len(maze[0])-1 && maze[i][j+1] == 0 {
			if res := dfs(i, j+1, append(path, "右")); res != nil {
				return res
			}
		}
		return nil
	}
	return dfs(start[0], start[1], []string{})
}
