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
	fmt.Println("Day 8, puzzle 1.\n\nImmediately before any instruction is executed a second time, what value is in the accumulator?")

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
	instructions := strings.Split(string(content), "\n")

	executed := make(map[int]bool)
	accumulator := 0
	programCounter := 0
	for true {
		opcode := instructions[programCounter][:3]
		amount, err := strconv.Atoi(instructions[programCounter][4:])
		if err != nil {
			log.Fatal(err)
		}
		if executed[programCounter] {
			fmt.Printf("Accumulator value is: %v\n", accumulator)
			break
		}

		switch opcode {
		case "nop":
			executed[programCounter] = true
			programCounter++
		case "acc":
			executed[programCounter] = true
			accumulator += amount
			programCounter++
		case "jmp":
			executed[programCounter] = true
			programCounter += amount
		default:
			log.Fatalln()

		}
	}
}

func puzzle2() {
	fmt.Println("Day 8, puzzle 2.\n\nWhat is the value of the accumulator after the program terminates?")

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
	instructions := strings.Split(string(content), "\n")

	executed := make(map[int]bool)
	accumulator := 0
	programCounter := 0
	instructionsCopy := make([]string, len(instructions))
	for true {
		if programCounter >= len(instructions) {
			fmt.Printf("PROGRAM ENDED.Accumulator value is: %v\n", accumulator)
			return
		}
		opcode := instructions[programCounter][:3]
		amount, err := strconv.Atoi(instructions[programCounter][4:])
		if err != nil {
			log.Fatal(err)
		}
		if executed[programCounter] {
			fmt.Printf("INFINITE LOOP DETECTED. Accumulator value is: %v\n", accumulator)
			return
		}

		switch opcode {
		case "nop":
			copy(instructionsCopy, instructions)
			if acc, ends := correctAndTry(instructionsCopy, programCounter, accumulator, cloneMap(executed)); ends {
				fmt.Printf("PROGRAM ENDED.Accumulator value is: %v\n", acc)
				return
			}
			executed[programCounter] = true
			programCounter++
		case "acc":
			executed[programCounter] = true
			accumulator += amount
			programCounter++
		case "jmp":
			copy(instructionsCopy, instructions)
			if acc, ends := correctAndTry(instructionsCopy, programCounter, accumulator, cloneMap(executed)); ends {
				fmt.Printf("PROGRAM ENDED.Accumulator value is: %v\n", acc)
				return
			}
			executed[programCounter] = true
			programCounter += amount
		default:
			log.Fatalln()

		}
	}
}

// findValidCode switches the current operation from jump to nop or viceversa and then
// tests if execution ends.
func correctAndTry(instructions []string, programCounter, accumulator int, executed map[int]bool) (int, bool) {

	switch instructions[programCounter][:3] {
	case "nop":
		instructions[programCounter] = strings.Replace(instructions[programCounter], "nop", "jmp", 1)
	case "jmp":
		instructions[programCounter] = strings.Replace(instructions[programCounter], "jmp", "nop", 1)
	default:
		log.Fatalln()

	}

	for true {
		if programCounter >= len(instructions) {
			return accumulator, true
		}
		opcode := instructions[programCounter][:3]
		amount, err := strconv.Atoi(instructions[programCounter][4:])
		if err != nil {
			log.Fatal(err)
		}
		if executed[programCounter] {
			return accumulator, false
		}

		executed[programCounter] = true
		switch opcode {
		case "nop":
			programCounter++
		case "acc":
			accumulator += amount
			programCounter++
		case "jmp":
			programCounter += amount
		default:
			log.Fatalln()

		}
	}

	return 0, false
}

func cloneMap(toClone map[int]bool) map[int]bool {

	res := make(map[int]bool)
	for key, value := range toClone {
		res[key] = value
	}

	return res
}
