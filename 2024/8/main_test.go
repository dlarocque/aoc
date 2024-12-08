package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`
	exampleSolution := 34

	res := solve(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}