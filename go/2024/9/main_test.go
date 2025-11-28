package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := `2333133121414131402`
	exampleSolution := 2858

	res := solve(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}
