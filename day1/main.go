package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	Question1()
	Question2()

}

func Question1() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	oldMax := 0
	counter := 0
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			if counter > oldMax {
				oldMax = counter
			}
			counter = 0

		} else {
			v, _ := strconv.Atoi(l)
			counter += v
		}
	}
	fmt.Println("Q1: ", oldMax)
}

func Question2() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	elfBags := make([]int, 0)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			elfBags = append(elfBags, counter)
			counter = 0
		} else {
			v, _ := strconv.Atoi(l)
			counter += v
		}
	}

	sort.Slice(elfBags, func(i, j int) bool {
		return elfBags[i] > elfBags[j]
	})

	fmt.Println("Q2: ", elfBags[0]+elfBags[1]+elfBags[2])
}
