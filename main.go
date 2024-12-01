package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: setup <year> <day>")
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Sprintf("Failed to parse argument year: %s", err))
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(fmt.Sprintf("Failed to parse argument year: %s", err))
	}

	currYear, currDay := time.Now().Year(), time.Now().Day()
	if year < 2015 || year > currYear {
		log.Fatalf("year is out of range [2015,%d]: %d", currYear, year)
	}

	if day < 1 || day > currDay+1 { // Using the current day would cause issues if we're in an early timezone.
		log.Fatalf("day is out of range [1,%d]: %d", currDay+1, day)
	}

	input := getInput(year, day)
	solutionDirName := fmt.Sprintf("%d/%d", year, day)
	err = os.MkdirAll(solutionDirName, 0750)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(solutionDirName+"/input.txt", input, 0660)
	if err != nil {
		log.Fatal(err)
	}

	// generate solution template
}

func getInput(year int, day int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("failed to create http request: %v", err)
	}
	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", getCookie()
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("failed to make client request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	return body
}

func getCookie() string {
	val, exists := os.LookupEnv("AOC_SESSION_COOKIE")
	if !exists {
		panic("Session cookie could not be found in AOC_SESSION_COOKIE environment variable.")
	}
	return val
}
