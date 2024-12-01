package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := `3   4
4   3
2   5
1   3
3   9
3   3`
	exampleSolution := 31

	res := solve2(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}
