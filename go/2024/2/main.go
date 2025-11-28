package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(solve(string(data)))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func strArrToInt(arr []string) []int {
	res := make([]int, len(arr))
	for i, v := range arr {
		conv, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		res[i] = conv
	}

	return res
}

func solve(input string) int {
	lines := strings.Split(input, "\n")

	numSafe := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		reports := strArrToInt(strings.Split(line, " "))
		if isSafe(reports) || isSafeWithDeletions(reports) {
			numSafe++
		}
	}

	return numSafe
}

func isSafe(reports []int) bool {
	hasInc := false
	hasDec := false

	for i := 1; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]
		if diff > 0 {
			hasInc = true
		} else {
			hasDec = true
		}

		if (hasInc && hasDec) || abs(diff) > 3 || diff == 0 {
			return false
		}
	}
	return true
}

// brute force. cba to think of all edge cases for fast algo.
func isSafeWithDeletions(reports []int) bool {
	for i := 0; i < len(reports); i++ {
		reportsWithDeletion := make([]int, len(reports))
		copy(reportsWithDeletion, reports)
		if i == 0 {
			reportsWithDeletion = reportsWithDeletion[i+1:]
		} else if i == len(reports)-1 {
			reportsWithDeletion = reportsWithDeletion[:i]
		} else {
			reportsWithDeletion = append(reportsWithDeletion[:i], reportsWithDeletion[i+1:]...)
		}

		if isSafe(reportsWithDeletion) {
			return true
		}

	}

	return false
}
