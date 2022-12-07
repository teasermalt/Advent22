package main

import (
	"fmt"
	"os"
	"bufio"
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

func calculate_score(list []string) int{
	score := 0

	for i := 0; i < len(list); i++ {
		opponent := list[i][0:1]
		player := list[i][2:3]

		rock := 1 //A, X, lose
		paper := 2 //B, Y, draw
		scissors := 3 //C, Z, win

		draw := 3
		win := 6
		
		
		if opponent == "A" && player == "Y" {
			score = score + draw + rock
		}
		if opponent == "A" && player == "Z" {
			score = score + win + paper
		}
		if opponent == "A" && player == "X" {
			score = score + scissors
		}
		if opponent == "B" && player == "X" {
			score = score + rock
		}
		if opponent == "B" && player == "Z" {
			score = score + win + scissors
		}
		if opponent == "B" && player == "Y" {
			score = score + draw + paper
		}
		if opponent == "C" && player == "X" {
			score = score + paper
		}
		if opponent == "C" && player == "Y" {
			score = score + draw + scissors
		}
		if opponent == "C" && player == "Z" {
			score = score + win + rock
		}


		//fmt.Println(opponent, player)
	}
	
	return score
}

func main() {

	//read input from file and store in a slice
	input := ReadInput()

	//calculate score
	
	fmt.Println(calculate_score(input))
}