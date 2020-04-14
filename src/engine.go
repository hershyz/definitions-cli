package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func run(term string) {

	currentfile, err := os.Open(filepath)

	if err != nil {
		log.Fatal("failed to open file: " + filepath)
	}

	fileScanner := bufio.NewScanner(currentfile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	currentfile.Close()

	var i int = 0
	var found bool = false

	for i < len(lines) {
		if strings.Contains(lines[i], term) {
			if containsTrigger(lines[i]) {
				found = true
				fmt.Println("line: ", i+1, " -", lines[i])
			}
		}
		i++
	}

	if !found {
		fmt.Println("no definitions found for: ", term)
	}
}

func containsTrigger(line string) bool {

	var i int = 0
	for i < len(triggers) {
		if strings.Contains(line, triggers[i]) {
			return true
		}
		i++
	}

	return false
}
