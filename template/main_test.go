package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := ``
	exampleSolution := 0

	res := solve(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}
