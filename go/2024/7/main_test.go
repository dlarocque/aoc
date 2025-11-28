package main

import "testing"

func TestExample(t *testing.T) {
	exampleInput := `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`
	exampleSolution := 11387

	res := solve(exampleInput)
	if res != exampleSolution {
		t.Errorf("\nactual: \n%d\n\nexpected:\n%d", res, exampleSolution)
	}
}
