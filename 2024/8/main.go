package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
)

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
	g := make([][]string, len(lines))
	for i, line := range lines {
		g[i] = strings.Split(line, "")
	}

	antennas := make(map[string][][]int)
	antinodes := make([][]bool, len(g))
	for i := range antinodes {
		antinodes[i] = make([]bool, len(g[i]))
	}	

	for i := range g {
		for j := range g[i] {
			if g[i][j] != "." {
				antennas[g[i][j]] = append(antennas[g[i][j]], []int{i, j})
			}
		}
	}

	res := 0
	for _, locs := range antennas {
		for j, aPos := range locs {
			for k, bPos := range locs {
				if j == k {
					continue
				}

				dist := []int {bPos[0] - aPos[0], bPos[1] - aPos[1]}
				antinodePos := []int {bPos[0], bPos[1]}
				for antinodePos[0] >= 0 && antinodePos[0] < len(antinodes) && antinodePos[1] >= 0 && antinodePos[1] < len(antinodes[0]) {
					if !antinodes[antinodePos[0]][antinodePos[1]] {
						res++
						antinodes[antinodePos[0]][antinodePos[1]] = true
					}
					
					antinodePos[0] += dist[0]
					antinodePos[1] += dist[1]
				}
			}	
		}
	}
	
	return res
}
