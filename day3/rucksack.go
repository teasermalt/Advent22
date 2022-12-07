package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() []string  {
	//read input from file args[1]
	filename := os.Args[1]
	list := make([]string, 0)

	//open file
	file, err := os.Open(filename)
	if err != nil {

		fmt.Println("Error opening file")
		os.Exit(1)
	}

	//add each line to slice
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	//close file
	defer file.Close()
	//fmt.Println(len(list))
	return list
}

func parse_list(list []string) {
	//return map of priorities
	//lowercase := "abcdefghijklmnopqrstuvwxyz"
	//uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//priorities := make(map[string]int)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}

func main() {
	score := 0
	list := ReadInput()
	fmt.Println(list, score)
}