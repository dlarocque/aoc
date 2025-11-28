package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	exampleSolution := 9

	res := solve2(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}
