package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Round struct {
	Opp string
	Me  string
}

var rounds = make([]Round, 0)

type Score map[Round]int

var scores = make(Score)
var scoresP2 = make(Score)

func main() {
	loadInput()
	Question1()
	Question2()
}

func loadInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), " ")
		rounds = append(rounds, Round{splitText[0], splitText[1]})
	}

	scores[Round{"A", "X"}] = 4
	scores[Round{"B", "X"}] = 1
	scores[Round{"C", "X"}] = 7
	scores[Round{"A", "Y"}] = 8
	scores[Round{"B", "Y"}] = 5
	scores[Round{"C", "Y"}] = 2
	scores[Round{"A", "Z"}] = 3
	scores[Round{"B", "Z"}] = 9
	scores[Round{"C", "Z"}] = 6

	scoresP2[Round{"A", "X"}] = 3
	scoresP2[Round{"B", "X"}] = 1
	scoresP2[Round{"C", "X"}] = 2
	scoresP2[Round{"A", "Y"}] = 4
	scoresP2[Round{"B", "Y"}] = 5
	scoresP2[Round{"C", "Y"}] = 6
	scoresP2[Round{"A", "Z"}] = 8
	scoresP2[Round{"B", "Z"}] = 9
	scoresP2[Round{"C", "Z"}] = 7

}

func Question1() {
	score := 0
	for _, v := range rounds {
		score += scores[v]
	}

	fmt.Println("Question 1: ", score)
}

func Question2() {
	score := 0
	for _, v := range rounds {
		score += scoresP2[v]
	}

	fmt.Println("Question 2: ", score)
}
