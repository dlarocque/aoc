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

func solve(input string) int {
	arr := strings.Split(input, "")

	// format disk string
	disk := make([]int, 0)
	r := len(arr) - 1
	if r%2 == 1 {
		r--
	}
	for l := range arr {
		if l%2 == 1 {
			freeSize, _ := strconv.Atoi(arr[l])
			fileSize, _ := strconv.Atoi(arr[r])
			for fileSize > freeSize && l <= r {
				r -= 2
				fileSize, _ = strconv.Atoi(arr[r])
			}

			rId := r / 2
			for i := 0; i < freeSize; i++ {
				if i < fileSize {
					disk = append(disk, rId)
				} else {
					disk = append(disk, -1)
				}
			}
			fmt.Printf("free space: r %d, rId: %d, freeSize: %d, fileSize: %d\n", r, rId, freeSize, fileSize)
			printDisk(disk)
			fmt.Println("00992111777.44.333....5555.6666.....8888..")

			r -= 2
		} else {
			fileSize, _ := strconv.Atoi(arr[l])
			id := l / 2
			for i := 0; i < fileSize; i++ {
				disk = append(disk, id)
			}
		}
	}

	fmt.Printf("input: %s\ndisk: %v\n", input, disk)

	res := 0
	r = len(disk) - 1
	for l := range disk {
		if l > r {
			break
		} else if disk[l] >= 0 {
			res += disk[l] * l
		} else {
			for disk[r] < 0 {
				r--
			}

			fmt.Printf("empty space. replacing with %d\n", disk[r])
			res += disk[r] * l
			r--
		}
		fmt.Printf("res: %d\n", res)
	}

	return res
}

func printDisk(disk []int) {
	s := ""
	for i := range disk {
		if disk[i] == -1 {
			s += "."
		} else {
			s += fmt.Sprint(disk[i])
		}
	}

	fmt.Println(s)
}
