package main

import (
	"fmt"
	lem "lem/methods"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("😞 ERROR: Usage EX: go run . example00.txt")
		return
	}

	// Read the file
	content, err := os.ReadFile("examples/" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	dataFile := strings.Split(string(content), "\n")
	if len(dataFile) == 0 {
		fmt.Println("😞 invalid file: the file is empty")
	}

	// Create an Input struct and populate it with the file data
	data := &lem.Input{}
	err = lem.GetData(data, dataFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a graph
	graph := &lem.Graph{}

	// Adding rooms
	for _, v := range data.Rooms {
		if err := lem.AddRoom(graph, v); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Adding tunnels (links) between rooms
	for _, v := range data.Links {
		if err := lem.AddEdge(graph, v[0], v[1]); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// BFS to find all paths
	allPaths := &lem.Allpaths{}
	allPaths.Paths = lem.BFS(graph, data.StartR, data.EndR)
	if len(allPaths.Paths) == 0 {
		fmt.Println(" 😞 There is no path between Start and End rooms")
		return
	}

	// Generate path combinations
	lem.GenerateCombinations(allPaths, graph)

	// Calculate Time for each combination
	time := lem.CalculateTravelTimesForEachCombination(allPaths, data.Ants)

	// Identify the optimal path
	lem.FindOptimalPath(allPaths, time, graph)
	lem.PrintAntMovement(allPaths, data.Ants)
}
