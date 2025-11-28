package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Map struct {
	m          [][]string
	x          int
	y          int
	dirX       int
	dirY       int
	numVisited int
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

	m := &Map{
		m:          graph,
		x:          startRow,
		y:          startCol,
		dirX:       -1,
		dirY:       0,
		numVisited: 1,
	}
	m.Print()

	m.Walk()
	return m.numVisited
}

func (m *Map) Walk() {
	nextX, nextY := m.x+m.dirX, m.y+m.dirY
	hitObstacle := false
	for !hitObstacle && 0 <= nextX && nextX < len(m.m) && 0 <= nextY && nextY < len(m.m[0]) {
		// fmt.Printf("x %d y %d, nextX %d nextY %d\n", m.x, m.y, nextX, nextY)
		if m.m[nextX][nextY] == "#" {
			hitObstacle = true
			// fmt.Println("HIT OBSTACLE")
		} else {
			// We haven't hit an obstacle, so we can move to the next position
			if m.m[m.x][m.y] != "X" {
				m.m[m.x][m.y] = "X"
				m.numVisited += 1
			}
			m.x = nextX
			m.y = nextY
			nextX += m.dirX
			nextY += m.dirY

			// Mark the position as visited, increasing the count if we haven't visited
		}

		m.Print()
	}

	// If the next character is an obstacle, we have to rotate 90 degrees to the right
	// -1, 0 -> 0, 1
	// 0, 1  -> 1, 0
	// 1, 0  -> 0, -1
	// 0, -1 -> -1, 0
	if hitObstacle {
		// There has to be a more clever way to do this
		if m.dirX == -1 && m.dirY == 0 {
			m.dirX = 0
			m.dirY = 1
		} else if m.dirX == 0 && m.dirY == 1 {
			m.dirX = 1
			m.dirY = 0
		} else if m.dirX == 1 && m.dirY == 0 {
			m.dirX = 0
			m.dirY = -1
		} else if m.dirX == 0 && m.dirY == -1 {
			m.dirX = -1
			m.dirY = 0
		} else {
			panic(fmt.Sprintf("invalid location dirX: %d, dirY: %d\n\n", m.dirX, m.dirY))
		}

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
