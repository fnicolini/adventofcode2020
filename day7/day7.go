package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/yourbasic/graph"
)

func main() {

	puzzle1()
	puzzle2()
}

func puzzle1() {
	fmt.Println("Day 7, puzzle 1.\n\nHow many bag colors can eventually contain at least one shiny gold bag?")

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
	rules := strings.Split(string(content), "\n")

	/*
		Assign to each color a number, because graph package does not support vertex labeling.
	*/
	bagIDs := make(map[string]int)
	bagContents := make(map[int][]string)
	for i, rule := range rules {
		currentBag := rule[:strings.Index(rule, "bags")-1]
		bagIDs[currentBag] = i
		bagContentsExp := regexp.MustCompile(`\d\s(.*?)\sbag`)
		subMatches := bagContentsExp.FindAllStringSubmatch(rule, -1)
		if subMatches != nil {
			for _, bag := range subMatches {
				bagContents[i] = append(bagContents[i], bag[1])
			}
		}
	}

	// Create a directed graph. The direction of the edge means "is contained in"
	g := graph.New(len(bagIDs))
	for k, v := range bagContents {
		for _, bag := range v {
			g.Add(bagIDs[bag], k)
		}
	}

	counter := 0
	// starting from the vertex of shiny gold, count every reachable node.
	graph.BFS(g, bagIDs["shiny gold"], func(v, w int, _ int64) {
		counter++
	})

	fmt.Printf("The number of possible bags is: %v.\n", counter)

}

func puzzle2() {
	fmt.Println("Day 7, puzzle 2.\n\nHow many individual bags are required inside your single shiny gold bag?")

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
	rules := strings.Split(string(content), "\n")

	/*
		Assign to each color a number, because graph package does not support vertex labeling.
	*/
	bagIDs := make(map[string]int)
	bagContents := make(map[int][]string)
	bagWeights := make(map[int][]int64)
	for i, rule := range rules {
		currentBag := rule[:strings.Index(rule, "bags")-1]
		bagIDs[currentBag] = i
		bagContentsExp := regexp.MustCompile(`(\d)\s(.*?)\sbag`)
		subMatches := bagContentsExp.FindAllStringSubmatch(rule, -1)
		//fmt.Printf("%q\n", subMatches)
		if subMatches != nil {
			for _, bag := range subMatches {
				bagContents[i] = append(bagContents[i], bag[2])
				x, err := strconv.Atoi(bag[1])
				if err != nil {
					log.Fatal(err)
				}
				bagWeights[i] = append(bagWeights[i], int64(x))
			}
		}
	}
	// Create a directed graph. The direction of the edge means "contains x number of bags"
	// where x is the weight of the graph
	g := graph.New(len(bagIDs))
	for k, v := range bagContents {
		for i, bag := range v {
			g.AddCost(k, bagIDs[bag], bagWeights[k][i])
		}
	}

	// starting from the vertex of shiny gold, count every reachable node.

	fmt.Printf("The number of required bags is: %v.\n", traverseDFS(g, bagIDs["shiny gold"]))

}

// traverseDFS executes a DFS search in the graph and calculates the number of bags
func traverseDFS(g graph.Iterator, v int) int64 {
	res := int64(0)
	g.Visit(v, func(w int, c int64) (_ bool) {
		res += c + (c * traverseDFS(g, w))
		return
	})
	return res
}
