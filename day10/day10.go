package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var neighbors map[int][]int
var adapterSequence []int
var adaptersChecked map[int]int

func main() {

	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println("Day 10, puzzle 1.\n\nWhat is the number of 1-jolt differences multiplied by the number of 3-jolt differences?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	// Remove last newline if present
	if rune(content[len(content)-1]) == '\n' {
		content = content[:len(content)-1]
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	strings := strings.Split(string(content), "\n")
	numbers := make([]int, len(strings))

	for i, s := range strings {
		numbers[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	numbers = append(numbers, 0)
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	diffs := make(map[int]int)
	for i := 0; i < len(numbers)-1; i++ {
		diffs[numbers[i+1]-numbers[i]]++
	}
	fmt.Printf("The answer is: %v", diffs[1]*diffs[3])
}

func puzzle2() {
	fmt.Println("\n\nDay 10, puzzle 2.\n\nWhat is the total number of distinct ways you can arrange the adapters to connect the charging outlet to your device?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	// Remove last newline if present
	if rune(content[len(content)-1]) == '\n' {
		content = content[:len(content)-1]
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	strings := strings.Split(string(content), "\n")
	adapterSequence = make([]int, len(strings))
	for i, s := range strings {
		adapterSequence[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
	}
	adapterSequence = append(adapterSequence, 0)
	sort.Ints(adapterSequence)
	adapterSequence = append(adapterSequence, adapterSequence[len(adapterSequence)-1]+3)
	neighbors = make(map[int][]int)

	adaptersChecked = make(map[int]int, len(adapterSequence))

	for i, n := range adapterSequence {
		for j := i + 1; j < len(adapterSequence); j++ {
			if adapterSequence[j]-adapterSequence[i] <= 3 {
				neighbors[n] = append(neighbors[n], adapterSequence[j])
			} else {
				break
			}
		}
	}

	//fmt.Println(numbers)
	fmt.Printf("The answer is %v\n", countSequence(0))
}

func countSequence(startAdapter int) int {

	counter := 0
	furthestPossibleAdapter := startAdapter + 3

	//fmt.Println(startAdapter)

	if furthestPossibleAdapter >= adapterSequence[len(adapterSequence)-1] {
		return 1
	} else if _, ok := adaptersChecked[startAdapter]; ok {
		return adaptersChecked[startAdapter]
	}
	for _, a := range neighbors[startAdapter] {
		counter += countSequence(a)
	}

	adaptersChecked[startAdapter] = counter

	return counter
}
