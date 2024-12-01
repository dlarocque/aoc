package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(solve1(string(data)))
	fmt.Println(solve2(string(data)))
}

func solve2(in string) int {
	lines := strings.Split(in, "\n")
	ll, rl := make([]int, len(lines)), make([]int, len(lines))
	for i, v := range lines {
		arr := strings.Split(v, "   ")
		if len(arr) != 2 {
			continue
		}
		l, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}

		r, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		ll[i] = l
		rl[i] = r
	}

	sort.Slice(ll, func(i, j int) bool { return ll[i] < ll[j] })
	sort.Slice(rl, func(i, j int) bool { return rl[i] < rl[j] })

	occur := make(map[int]int)
	res := 0
	j := 0
	fmt.Println(ll, rl)
	for i := 0; i < len(ll); i++ {
		_, ok := occur[ll[i]]
		if !ok {
			// find the first occurence of the number in the right list
			for j < len(rl) && ll[i] > rl[j] {
				fmt.Println(i, j)
				j++
			}
			fmt.Println(i, j)
			// iterate over all occurences of the number in the right list
			for j < len(rl) && ll[i] == rl[j] {
				occur[ll[i]] += ll[i]
				j++
			}
			fmt.Println(i, j)
		}

		res += occur[ll[i]]
	}

	fmt.Println(occur)
	return res
}

func solve1(in string) int {
	lines := strings.Split(in, "\n")
	ll, rl := make([]int, len(lines)), make([]int, len(lines))
	for i, v := range lines {
		arr := strings.Split(v, "   ")
		if len(arr) != 2 {
			continue
		}
		fmt.Println(arr)
		l, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}

		r, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		ll[i] = l
		rl[i] = r
	}

	sort.Slice(ll, func(i, j int) bool { return ll[i] < ll[j] })
	sort.Slice(rl, func(i, j int) bool { return rl[i] < rl[j] })
	fmt.Println(ll, rl)
	listLen := len(ll)
	res := 0
	for i := 0; i < listLen; i++ {
		diff := ll[i] - rl[i]
		if diff < 0 {
			diff = -diff
		}
		fmt.Println(ll[i], rl[i], diff)

		res += diff
	}

	return res
}
