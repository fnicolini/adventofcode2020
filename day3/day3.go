package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println("Day 3, puzzle 1.\n\n Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	lines := strings.Split(string(content), "\n")

	/*
		For each line, increment the counter if a tree is encountered.
	*/
	counter := 0
	for i, lineStr := range lines {

		line := []rune(lineStr)
		lineIndex := (i * 3) % (len(line))

		fmt.Println(lineIndex)
		if string(line[lineIndex]) == "#" {
			counter++
		}
	}
	fmt.Printf("I would encounter %v trees.\n", counter)

}

func puzzle2() {
	fmt.Println("\nDay 3, puzzle 2.\n\nList of slopes:\n" +
		"Right 1, down 1.\n" +
		"Right 3, down 1.\n" +
		"Right 5, down 1.\n" +
		"Right 7, down 1.\n" +
		"Right 1, down 2.\n" +
		"What do you get if you multiply together the number of trees encountered on each of the listed slopes?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	lines := strings.Split(string(content), "\n")

	/*
		For each line, increment the counters if a tree is encountered.
	*/
	var counters [5]uint
	l := len(lines[0])
	for i, lineStr := range lines {
		line := []rune(lineStr)
		slope1Index := i % l
		slope2Index := (i * 3) % l
		slope3Index := (i * 5) % l
		slope4Index := (i * 7) % l
		slope5Index := (i / 2) % l

		if string(line[slope1Index]) == "#" {
			counters[0]++
		}

		if string(line[slope2Index]) == "#" {
			counters[1]++
		}

		if string(line[slope3Index]) == "#" {
			counters[2]++
		}

		if string(line[slope4Index]) == "#" {
			counters[3]++
		}

		if i%2 == 0 && string(line[slope5Index]) == "#" {
			counters[4]++
		}
	}

	res := uint(1)
	for _, counter := range counters {
		res *= counter
	}
	fmt.Printf("I would encounter %v trees.\n", res)

}
