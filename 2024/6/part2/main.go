package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Map struct {
	m        [][]string
	x        int
	y        int
	numLoops int
	dirs     [][]int
	dir      int
}

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	result := solve(string(data))

	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Execution time: %d ms\n", elapsed.Milliseconds()) // Output in milliseconds
}

func solve(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	graph := make([][]string, len(lines))
	startRow, startCol := 0, 0
	for r, line := range lines {
		graph[r] = strings.Split(line, "")

		for c, x := range graph[r] {
			if x == "^" {
				startRow = r
				startCol = c
				graph[r][c] = "."
			}
		}
	}

	dirs := [][]int{
		[]int{-1, 0}, // up
		[]int{0, 1},  // right
		[]int{1, 0},  // down
		[]int{0, -1}, // left
	}

	m := &Map{
		m:        graph,
		x:        startRow,
		y:        startCol,
		numLoops: 0,
		dirs:     dirs,
		dir:      0,
	}
	// m.Print()

	m.Walk()
	return m.numLoops
}

func (m *Map) Walk() {
	nextX, nextY := m.x+m.dirs[m.dir][0], m.y+m.dirs[m.dir][1]
	hitObstacle := false
	for !hitObstacle && 0 <= nextX && nextX < len(m.m) && 0 <= nextY && nextY < len(m.m[0]) {
		if m.m[nextX][nextY] == "#" {
			hitObstacle = true
		} else {
			if m.m[m.x][m.y] == "." {
				m.m[m.x][m.y] = fmt.Sprint(m.dir)
			}

			m.x = nextX
			m.y = nextY

			mTemp := make([][]string, len(m.m))
			for i := range m.m {
				mTemp[i] = make([]string, len(m.m[i]))
				copy(mTemp[i], m.m[i])
			}
			blockedDir := (m.dir + 1) % 4
			currX, currY := m.x, m.y
			checkX, checkY := m.x+m.dirs[blockedDir][0], m.y+m.dirs[blockedDir][1]
			for 0 <= currX && currX < len(mTemp) && 0 <= currY && currY < len(mTemp[0]) {
				if mTemp[currX][currY] == fmt.Sprint(blockedDir) {
					m.numLoops++
					break
				} else if 0 <= checkX && checkX < len(mTemp) && 0 <= checkY && checkY < len(mTemp[0]) && mTemp[checkX][checkY] == "#" {
					blockedDir = (blockedDir + 1) % 4
				} else {
					mTemp[currX][currY] = fmt.Sprint(blockedDir)
					currX += m.dirs[blockedDir][0]
					currY += m.dirs[blockedDir][1]
				}

				checkX += m.dirs[blockedDir][0]
				checkY += m.dirs[blockedDir][1]
			}

			nextX += m.dirs[m.dir][0]
			nextY += m.dirs[m.dir][1]
		}
	}

	if hitObstacle {
		m.dir = (m.dir + 1) % 4
		m.Walk() // Keep going in the other direction!
	}

}

func (m *Map) Print() {
	graphStr := ""
	for x, row := range m.m {
		for y, elem := range row {
			if x == m.x && y == m.y {
				graphStr += "v"
			} else {
				graphStr += elem
			}
		}
		graphStr += "\n"
	}

	fmt.Println(graphStr)
}
