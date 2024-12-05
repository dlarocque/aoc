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

	fmt.Println("rules: ", rules)
	fmt.Println("updates: ", updates)

	ordering := make(map[int]map[int]bool, len(rules))
	for _, rule := range rules {
		nums := strings.Split(rule, "|")
		first, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		if ordering[first] == nil {
			ordering[first] = make(map[int]bool, len(rules))
		}
		ordering[first][second] = true
	}

	fmt.Println("ordering: ", ordering)

	res := 0
	for i := 0; i < len(updates); i++ {
		numsStrs := strings.Split(updates[i], ",")
		nums := make([]int, len(numsStrs))
		for i, numStr := range numsStrs {
			if numStr == "" {
				break
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums[i] = num
		}

		valid := true
		for i, num := range nums {
			if !isValid(ordering[num], nums[:i], num) {
				valid = false
			}
		}

		if !valid {
			var ordered []int
			for len(nums) > 0 {
				num := nums[0]
				nums = nums[1:]
				canPlace := true
				for _, other := range nums {
					if ordering[num][other] {
						canPlace = false
					}
				}

				if canPlace {
					ordered = append(ordered, num)
				} else {
					nums = append(nums, num)
				}
			}

			res += ordered[len(ordered)/2]
		}

	}

	return res
}

func isValid(notAllowed map[int]bool, preceeding []int, num int) bool {
	for _, num := range preceeding {
		_, ok := notAllowed[num]
		if ok {
			return false
		}

	}

	return true
}
