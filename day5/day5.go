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
	fmt.Println("Day 5, puzzle 1.\n\n As a sanity check, look through your list of boarding passes. " +
		"What is the highest seat ID on a boarding pass?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	boardinPassList := strings.Split(string(content), "\n")

	/*
		For each passport, increment the counter if it's valid.
	*/
	maxSeatID := int64(0)
	for _, boardingPass := range boardinPassList {
		// Trasnform to binary
		boardingPass := strings.ReplaceAll(boardingPass, "F", "0")
		boardingPass = strings.ReplaceAll(boardingPass, "B", "1")
		boardingPass = strings.ReplaceAll(boardingPass, "R", "1")
		boardingPass = strings.ReplaceAll(boardingPass, "L", "0")

		row, err := strconv.ParseInt(boardingPass[:len(boardingPass)-3], 2, 8)
		if err != nil {
			log.Fatal(err)
		}
		column, err := strconv.ParseInt(boardingPass[len(boardingPass)-3:], 2, 7)
		if err != nil {
			log.Fatal(err)
		}

		seatID := row*8 + column

		if seatID > maxSeatID {
			maxSeatID = seatID
		}

	}

	fmt.Printf("The highest seat ID is %v.\n", maxSeatID)

}

func puzzle2() {
	fmt.Println("Day 5, puzzle 2.\n\n What is the ID of your seat?")

	//Read file content in slice of bytes
	content, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	/*
		Cast slice of bytes to a string (which contains the whole file) and then
		use the split function to create a slice using a separator (newline)
	*/
	boardinPassList := strings.Split(string(content), "\n")

	var seats [128][8]bool
	for _, boardingPass := range boardinPassList {
		// Trasnform to binary
		boardingPass := strings.ReplaceAll(boardingPass, "F", "0")
		boardingPass = strings.ReplaceAll(boardingPass, "B", "1")
		boardingPass = strings.ReplaceAll(boardingPass, "R", "1")
		boardingPass = strings.ReplaceAll(boardingPass, "L", "0")

		row, err := strconv.ParseInt(boardingPass[:len(boardingPass)-3], 2, 8)
		if err != nil {
			log.Fatal(err)
		}
		column, err := strconv.ParseInt(boardingPass[len(boardingPass)-3:], 2, 7)
		if err != nil {
			log.Fatal(err)
		}

		seats[row][column] = true

	}

	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			if seats[i][j] == false {
				seatID := i*8 + j
				fmt.Printf("Row=%v, column=%v seatID=%v\n", i, j, seatID)
			}
		}
	}

}
