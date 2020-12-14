package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println("Day 2, puzzle 1.\n\n How many passwords are valid according to their policies?")

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
	lines := strings.Split(string(content), "\n")

	/*
		For each line, increment the counter if the password matches the policy.
	*/
	counter := 0
	for _, line := range lines {

		dashIndex := strings.Index(line, "-")
		firstSpaceIndex := strings.Index(line, " ")
		colonsIndex := strings.Index(line[firstSpaceIndex+1:], ":") + firstSpaceIndex + 1
		secondSpaceIndex := strings.Index(line[colonsIndex+1:], " ") + colonsIndex + 1

		lowerLimit, err := strconv.Atoi(line[:dashIndex])
		if err != nil {
			log.Fatal(err)
		}

		upperLimit, err := strconv.Atoi(line[dashIndex+1 : firstSpaceIndex])
		if err != nil {
			log.Fatal(err)
		}

		c := line[firstSpaceIndex+1 : colonsIndex]
		password := line[secondSpaceIndex+1:]
		occurrences := strings.Count(password, c)

		if occurrences >= lowerLimit && occurrences <= upperLimit {
			counter++
		}
	}

	fmt.Printf("There are %v valid passwords according to their policies\n", counter)

}

func puzzle2() {
	fmt.Println("\nDay 2, puzzle 2.\n\n How many passwords are valid according to their policies?")

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
	lines := strings.Split(string(content), "\n")

	/*
		For each line, increment the validPasswords counter if the password matches the policy.
	*/
	validPasswords := 0
	for _, line := range lines {

		dashIndex := strings.Index(line, "-")
		firstSpaceIndex := strings.Index(line, " ")
		colonsIndex := strings.Index(line[firstSpaceIndex+1:], ":") + firstSpaceIndex + 1
		secondSpaceIndex := strings.Index(line[colonsIndex+1:], " ") + colonsIndex + 1

		firstPosition, err := strconv.Atoi(line[:dashIndex])
		if err != nil {
			log.Fatal(err)
		}

		secondPosition, err := strconv.Atoi(line[dashIndex+1 : firstSpaceIndex])
		if err != nil {
			log.Fatal(err)
		}

		c := line[firstSpaceIndex+1 : colonsIndex]
		password := line[secondSpaceIndex+1:]
		correctPositionCounter := 0

		if password[firstPosition-1:firstPosition] == c {
			correctPositionCounter++
		}

		if password[secondPosition-1:secondPosition] == c {
			correctPositionCounter++
		}

		if correctPositionCounter == 1 {
			validPasswords++
		}

	}

	fmt.Printf("There are %v valid passwords according to their policies\n", validPasswords)

}
