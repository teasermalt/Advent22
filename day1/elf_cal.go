package main

import (
	"fmt"
	"os"
	"bufio"
	//"io"

)

func readLines() ([]string, error) {
	filename := os.Args[1]
	list := make([]string, 0)
	
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	defer file.Close()
	return list, scanner.Err()

}

func main() {
	calories := readLines()
	fmt.Println(calories)
}