package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

type Equation struct {
	sol  int
	nums []int
}

func solve(input string) int {
	equations := processInput(input)
	res := 0

	for _, eq := range equations {
		prevSums := make([]int, 0)
		for _, num := range eq.nums {
			if len(prevSums) == 0 {
				prevSums = append(prevSums, num)
				continue
			}
			newSums := make([]int, 0)
			for _, prevSum := range prevSums {
				add := prevSum + num
				mult := prevSum * num

				concat, _ := strconv.Atoi(fmt.Sprintf("%d%d", prevSum, num))

				if add <= eq.sol {
					newSums = append(newSums, add)
				}
				if mult <= eq.sol {
					newSums = append(newSums, mult)
				}
				if concat <= eq.sol {
					newSums = append(newSums, concat)
				}
			}

			prevSums = newSums
		}
		for _, sum := range prevSums {
			if sum == eq.sol {
				res += eq.sol
				break
			}
		}
	}

	return res
}

func processInput(input string) []Equation {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // last line is empty
	equations := make([]Equation, len(lines))
	for i, line := range lines {
		if line == "" {
			continue
		}
		eq := strings.Split(line, ":")
		left, right := eq[0], eq[1]
		sol, _ := strconv.Atoi(left)
		rightArr := strings.Split(strings.Trim(right, " "), " ")
		nums := make([]int, len(rightArr))
		for j, elem := range rightArr {
			num, _ := strconv.Atoi(elem)
			nums[j] = num
		}

		equations[i] = Equation{
			sol:  sol,
			nums: nums,
		}
	}

	return equations
}
