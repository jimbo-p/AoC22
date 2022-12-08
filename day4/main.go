package main

import (
	"bufio"
	"fmt"
	"os"
)

type Compartment struct {
	Left  string
	Right string
}

type Group struct {
	e1 string
	e2 string
	e3 string
}

var list = make([]Compartment, 0)
var listQ2 = make([]Group, 0)

var badVals = make([]string, 0)
var prioSum int
var badgeSum int

func main() {
	loadInput()
	Question1()
	Question2()
}

func loadInput() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	g := Group{}
	for scanner.Scan() {
		s := scanner.Text()
		list = append(list, Compartment{s[:len(s)/2], s[len(s)/2:]})

		switch num := counter % 3; num {
		case 0:
			g.e1 = s
		case 1:
			g.e2 = s
		case 2:
			g.e3 = s
			listQ2 = append(listQ2, g)
			g = Group{}
		}
		counter += 1
	}
}

func Question1() {
	for _, v := range list {
		check := make(map[string]bool)
		for _, v := range v.Left {
			check[string(v)] = true
		}

		for _, v := range v.Right {
			if _, ok := check[string(v)]; ok {
				if casing := int(v); casing > 90 {
					prioSum += casing - 96
				} else {
					prioSum += casing - 38
				}
				badVals = append(badVals, string(v))
				break
			}
		}
	}

	fmt.Println("Question 1: ", prioSum)

}

func Question2() {
	for _, v := range listQ2 {
		check := make(map[string]int)
		for _, v := range v.e1 {
			check[string(v)] = 1
		}

		for _, v := range v.e2 {
			if val := check[string(v)]; val == 1 {
				check[string(v)] = 2
			}
		}

		for _, v := range v.e3 {
			if val := check[string(v)]; val == 2 {
				check[string(v)] = 3
			}
		}

		for k, v := range check {
			if v != 3 {
				delete(check, k)
			}
		}

		for k := range check {
			if int(k[0]) > 90 {
				badgeSum += int(k[0]) - 96
			} else {
				badgeSum += int(k[0]) - 38
			}
		}
	}
	fmt.Println("Question 2: ", badgeSum)
}
