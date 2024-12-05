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

func solve(input string) int {
	s := strings.Split(input, "\n\n")
	rules := strings.Split(s[0], "\n")
	updates := strings.Split(s[1], "\n")

	ordering := make(map[int]map[int]bool, len(rules))
	for _, rule := range rules {
		nums := strings.Split(rule, "|")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[1])

		if ordering[first] == nil {
			ordering[first] = make(map[int]bool)
		}
		ordering[first][second] = true
	}

	res := 0
	for _, update := range updates {
		if update == "" {
			continue
		}

		numsStr := strings.Split(update, ",")
		nums := make([]int, len(numsStr))
		valid := true

		for i, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums[i] = num

			if valid {
				for j := 0; j < i; j++ {
					if ordering[num][nums[j]] {
						valid = false
						break
					}
				}
			}
		}

		if !valid {
			res += reorderAndGetMedian(nums, ordering)
		}
	}

	return res
}

func reorderAndGetMedian(nums []int, ordering map[int]map[int]bool) int {
	n := len(nums)
	var ordered []int
	unprocessed := make(map[int]bool, n)
	for _, num := range nums {
		unprocessed[num] = true
	}

	for len(unprocessed) > 0 {
		for num := range unprocessed {
			canPlace := true
			for other := range unprocessed {
				if ordering[num][other] {
					canPlace = false
					break
				}
			}
			if canPlace {
				ordered = append(ordered, num)
				delete(unprocessed, num)
			}
		}
	}

	return ordered[len(ordered)/2]
}
