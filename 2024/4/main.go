package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solve2(string(data)))
}

func solve2(input string) int {
	// create a graph from the input
	inputLines := strings.Split(input, "\n")
	g := make([][]string, len(inputLines))
	for i, line := range inputLines {
		g[i] = strings.Split(strings.TrimSpace(line), "")
	}

	cnt := 0
	for i := 1; i < len(g)-1; i++ {
		for j := 1; j < len(g[i])-1; j++ {
			if g[i][j] == "A" {
				// from each corner, look for a diagonal "SAM" or "MAS"
				if ((g[i-1][j-1] == "S" && g[i+1][j+1] == "M") || (g[i-1][j-1] == "M" && g[i+1][j+1] == "S")) &&
					((g[i-1][j+1] == "S" && g[i+1][j-1] == "M") || (g[i-1][j+1] == "M" && g[i+1][j-1] == "S")) {
					cnt++
				}
			}
		}
	}

	return cnt
}

/*
func solve(input string) int {
	// create a graph from the input
	inputLines := strings.Split(input, "\n")
	g := make([][]string, len(inputLines))
	for i, line := range inputLines {
		chars := strings.Split(strings.TrimSpace(line), "")
		g[i] = make([]string, len(chars))
		g[i] = chars
	}

	cnt := 0
	for i := range g {
		for j := range g[i] {
			if g[i][j] == "X" {
				for iStep := -1; iStep <= 1; iStep++ {
					for jStep := -1; jStep <= 1; jStep++ {
						cnt += dfs(g, i, j, iStep, jStep, "M")
					}
				}
			}
		}
	}

	return cnt
}

func dfs(g [][]string, i, j, iStep, jStep int, next string) int {
	if i+iStep >= len(g) || i+iStep < 0 || j+jStep >= len(g[i+iStep]) || j+jStep < 0 {
		return 0
	}
	if g[i+iStep][j+jStep] == next {
		// finished the word!
		if next == "S" {
			return 1
		}

		// continue in the same direction for the next letter
		var nextI, nextJ int
		if iStep > 0 {
			nextI++
		} else if iStep < 0 {
			nextI--
		} else {
			nextI = 0
		}

		if jStep > 0 {
			nextJ++
		} else if jStep < 0 {
			nextJ--
		} else {
			nextJ = 0
		}

		var nextLetter string
		if next == "M" {
			nextLetter = "A"
		} else {
			nextLetter = "S"
		}

		return dfs(g, i, j, iStep+nextI, jStep+nextJ, nextLetter)
	}

	return 0
}
*/
