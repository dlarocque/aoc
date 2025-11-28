package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`(mul\(([0-9]{1,3},[0-9]{1,3})\))|((don't\(\))|(do\(\)))`)
	matches := re.FindAll([]byte(input), -1)
	if matches == nil {
		panic("found no matches in regexp")
	}

	res := 0
	enabled := true
	for _, match := range matches {
		matchStr := string(match)
		if matchStr == "don't()" {
			enabled = false
		} else if matchStr == "do()" {
			enabled = true
		} else if enabled {
			nums := strings.Split(string(match[4:len(match)-1]), ",")
			fmt.Println("nums", nums)
			a, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}

			b, err := strconv.Atoi((nums[1]))
			if err != nil {
				panic(err)
			}

			fmt.Printf("match: %s, a: %d, b: %d\n", match, a, b)
			res += a * b
		}
	}
	return res
}
