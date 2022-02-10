package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

type Maze [][]int

func (p point) add(a point) point {
	return point{p.i + a.i, p.j + a.j}
}

// ok 判断p点结合Maze步子记录中的情况判断是否可走
// true表示p点可走
func (maze Maze) ok(steps Maze, p point, start point) bool {
	// 是否在迷宫中
	rowsEdge, colsEdge := len(maze)-1, len(maze[0])-1
	if !(p.i >= 0 && p.j >= 0 && p.i <= rowsEdge && p.j <= colsEdge) {
		return false
	}
	// 是否有墙
	if maze[p.i][p.j] == 1 {
		return false
	}
	// 判断是否走过
	if steps[p.i][p.j] != -1 {
		return false
	}

	return true
}

func getMazeInfo() [][]int {
	file, err := os.Open("maze.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

var dirs = [4]point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func walk(maze Maze, start point, end point) ([][]int, int) {
	steps := make(Maze, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[i]))
		for j := range steps[i] {
			steps[i][j] = -1
		}
	}
	Q := []point{start}
	steps[start.i][start.j] = 0

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			// 判断可行性
			ok := maze.ok(steps, next, start)

			if !ok {
				continue
			}
			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			Q = append(Q, point{next.i, next.j})

			if next.i == end.i && next.j == end.j {
				return steps, steps[end.i][end.j]
			}
		}
	}
	return steps, steps[end.i][end.j]
}

func main() {
	maze := getMazeInfo()
	steps, minStep := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Printf("%v %d", steps, minStep)
}
