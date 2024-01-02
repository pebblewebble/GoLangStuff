package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

// REFERENCES
// https://www.baeldung.com/cs/maze-generation#:~:text=Another%20way%20of%20generating%20a,as%20visited%2C%20and%20so%20on.
// https://medium.com/swlh/fun-with-python-1-maze-generator-931639b4fb7e

func main() {
	width := 11
	height := 11
	// This program was designed to have and odd width and height
	for i := 0; i < 1; i++ {
		initializeMap := initialize(width, height)
		generatedMaze := generateMaze(initializeMap, width, height)
		visualizeMaze(generatedMaze)
	}
}

func initialize(height, width int) [][]string {
	//Creating the matrix
	gameMap := make([][]string, height)
	for heightIndex := 0; heightIndex < height; heightIndex++ {
		gameMap[heightIndex] = make([]string, width)
		for widthIndex := 0; widthIndex < width; widthIndex++ {
			//Generate outer walls and cells
			if heightIndex == 0 || heightIndex == height-1 || widthIndex == 0 || widthIndex == width-1 || heightIndex%2 == 0 || widthIndex%2 == 0 {
				gameMap[heightIndex][widthIndex] = color.HiBlueString("*")
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
	//Generates a random odd X,Y starting position but instead of constantly randomly generating a random position each time,
	//I could just minus 1 from the position that is not odd.
	startingPosition := generateStart(height, width)
	for startingPosition[0]%2 == 0 || startingPosition[1]%2 == 0 {
		startingPosition = generateStart(height, width)
	}
	fmt.Println(startingPosition)
	gameMap[startingPosition[0]][startingPosition[1]] = "X"

	//Able to choose what type of algo in the future
	gameMap = randomAlgo(gameMap, startingPosition, height, width)

	// walls := [][]int{}
	// walls = append()
	return gameMap
}

func randomAlgo(gameMap [][]string, startingPosition []int, height, width int) [][]string {
	//Move and remove the wall between each other
	//Randomize whether to go vertical or horizontal
	if rand.Intn(2) == 1 {
		//Vertical
		fmt.Println("VERTICAL")
		if rand.Intn(2) == 1 {
			//Going Up
			//If it does not go over the maze
			if startingPosition[0]-2 > 0 {
				fmt.Println("UP")
				gameMap[startingPosition[0]-1][startingPosition[1]] = " "
			}
		} else {
			//Going Down
			if startingPosition[0]+2 < height {
				fmt.Println("DOWN")
				gameMap[startingPosition[0]+1][startingPosition[1]] = " "
			}
		}
	} else {
		//Horizontal
		//If it does not go over the maze
		fmt.Println("HORIZONTAL")
		if rand.Intn(2) == 1 {
			if startingPosition[1]+2 < width {
				fmt.Println("RIGHT")
				gameMap[startingPosition[0]][startingPosition[1]+1] = " "
			}
		} else {
			if startingPosition[1]-2 > 0 {
				fmt.Println("LEFT")
				gameMap[startingPosition[0]][startingPosition[1]-1] = " "
			}
		}

	}
	return gameMap
}

func visualizeMaze(maze [][]string) {
	for i := range maze {
		fmt.Println(maze[i])
	}
}
