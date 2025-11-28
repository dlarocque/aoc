package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	err = copyDirectory("template", solutionDirName)
	if err != nil {
		panic(err)
	}
}

func copyDirectory(srcDir, targetDir string) error {
	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			if err := os.MkdirAll(destPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err := copyFile(path, destPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dest, srcInfo.Mode())
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
