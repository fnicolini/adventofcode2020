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
	fmt.Println("Day 6, puzzle 1.\n\n For each group, count the number of questions to which anyone answered \"yes\"." +
		" What is the sum of those counts?")

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
	groups := strings.Split(string(content), "\n\n")

	/*
		For each passport, increment the counter if it's valid.
	*/
	counter := 0
	for _, group := range groups {
		answers := make(map[rune]bool)
		people := strings.Split(string(group), "\n")
		for _, person := range people {
			str := []rune(person)
			for _, letter := range str {
				answers[letter] = true
			}
		}
		counter += len(answers)
	}

	fmt.Printf("The sum is %v.\n", counter)

}

func puzzle2() {
	fmt.Println("\nDay 6, puzzle 2.\n\n For each group, count the number of questions to which everyone answered \"yes\"." +
		" What is the sum of those counts?")

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
	groups := strings.Split(string(content), "\n\n")

	/*
		For each passport, increment the counter if it's valid.
	*/
	counter := 0
	for _, group := range groups {
		answers := make(map[rune]int)
		people := strings.Split(string(group), "\n")
		for _, person := range people {
			str := []rune(person)
			for _, letter := range str {
				answers[letter]++
			}
		}
		for _, element := range answers {
			if element == len(people) {
				counter++
			}
		}
	}

	fmt.Printf("The sum is %v.\n", counter)

}
