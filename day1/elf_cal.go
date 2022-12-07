package main

import (
	"fmt"
	"os"
	"bufio"
	//"io"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Println(readLines("elf_cal.txt"))
}