package main

import (
	"fmt"
	"math/rand"
)

// REFERENCES
// https://www.baeldung.com/cs/maze-generation#:~:text=Another%20way%20of%20generating%20a,as%20visited%2C%20and%20so%20on.
// https://medium.com/swlh/fun-with-python-1-maze-generator-931639b4fb7e

func main() {
	width := 21
	height := 21
	// This program was designed to have and odd width and height
	for i := 0; i < 10; i++ {
		initializeMap := initialize(width, height)
		generatedMaze := generateMaze(initializeMap, width, height)
		visualizeMaze(generatedMaze)
	}
}

func initialize(height, width int) [][]string {
	gameMap := make([][]string, height)
	for heightIndex := 0; heightIndex < height; heightIndex++ {
		gameMap[heightIndex] = make([]string, width)
		for widthIndex := 0; widthIndex < width; widthIndex++ {
			if heightIndex == 0 || heightIndex == height-1 || widthIndex == 0 || widthIndex == width-1 || heightIndex%2 == 0 || widthIndex%2 == 0 {
				gameMap[heightIndex][widthIndex] = "*"
			} else {
				gameMap[heightIndex][widthIndex] = " "
			}
		}
	}

	return gameMap
}

func generateStart(height, width int) []int {
	startingPosition := []int{rand.Intn(height-1-1) + 1, rand.Intn(width-2) + 1}
	return startingPosition
}

func generateMaze(gameMap [][]string, height, width int) [][]string {
	//
	startingPosition := generateStart(height, width)
	for startingPosition[0]%2 == 0 || startingPosition[1]%2 == 0 {
		startingPosition = generateStart(height, width)
	}
	fmt.Println(startingPosition)
	gameMap[startingPosition[0]][startingPosition[1]] = "X"
	// walls := [][]int{}
	// walls = append(walls)
	return gameMap
}

func visualizeMaze(maze [][]string) {
	for i := range maze {
		fmt.Println(maze[i])
	}
}
