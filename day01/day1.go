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

	fmt.Println("Day 1, puzzle 1.\n\n Find the two entries that sum to 2020; what do you get if you multiply them together?")

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
		For each element x of index i of the slice, loop from index i + 1 and check every
		 sum. Once the correct pair is found, print the result and exit the program.
	*/
	for i, a := range lines {
		for _, b := range lines[i+1:] {
			x, err := strconv.Atoi(a)
			if err != nil {
				log.Fatal(err)
			}

			y, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}

			if (x + y) == 2020 {
				z := x * y

				fmt.Printf("The two entries that sum to 2020 are %v and %v. Multiplying them together produces %v * %v = %v\n", x, y, x, y, z)

				return
			}

		}
	}

}

func puzzle2() {

	fmt.Println("\nDay 1, puzzle 2.\n\n What is the product of the three entries that sum to 2020?")

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
		For each element x of index i of the slice, loop from index i + 1 and check every
		 sum. Once the correct pair is found, print the result and exit the program.
	*/
	for i, a := range lines {
		for j, b := range lines[i+1:] {
			for _, c := range lines[j+1:] {

				x, err := strconv.Atoi(a)
				if err != nil {
					log.Fatal(err)
				}

				y, err := strconv.Atoi(b)
				if err != nil {
					log.Fatal(err)
				}

				z, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}

				if (x + y + z) == 2020 {
					p := x * y * z
					fmt.Printf("The three entries that sum to 2020 are %v, %v and %v. Multiplying them together produces %v * %v * %v = %v\n", x, y, z, x, y, z, p)

					return
				}
			}
		}
	}
}
