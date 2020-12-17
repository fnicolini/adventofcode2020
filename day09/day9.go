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
	fmt.Println("Day 10, puzzle 1.\n\nWhat is the first number that does not have this property?")

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

	for i := 25; i < len(numbers); i++ {
		if !findAddends(numbers[i-25:i], numbers[i]) {
			fmt.Printf("The first number to not have this property is: %v\n", numbers[i])
			return
		}
	}
}

func findAddends(x []int, n int) bool {
	for i, a := range x {
		for _, b := range x[i+1:] {
			if (a + b) == n {
				return true
			}
		}
	}
	return false
}

func puzzle2() {
	fmt.Println("\nDay 9, puzzle 2.\n\nWhat is the encryption weakness in your XMAS-encrypted list of numbers?")

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
	for i := 2; i < len(numbers); i++ {
		counter := 0
		for j := 0; j < len(numbers)-i; j++ {
			subgroup := numbers[j : j+i]
			if sumSlice(subgroup) == 3199139634 {
				fmt.Printf("Encryption weakness is: %v\n", minSlice(subgroup)+maxSlice(subgroup))
				return
			}
			counter++
		}
	}
}

func sumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func minSlice(s []int) int {
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func maxSlice(s []int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
