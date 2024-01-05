package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// REFERENCES
// https://www.baeldung.com/cs/maze-generation#:~:text=Another%20way%20of%20generating%20a,as%20visited%2C%20and%20so%20on.
// https://medium.com/swlh/fun-with-python-1-maze-generator-931639b4fb7e

func main() {
	width := 31
	height := 31
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
	//Empty cells is " ", so find number of " " and we can know the amount of empty cells and then while there is still empty cells
	//we continue the movement?
	//Need to improve movement so that if it rolls a direction but can't meet its inner condition, then redo another random direction again
	//Add condition to see if the direction being move has been visited b4.
	//A unvisited cell would be a cell with all the walls intact?
	emptyCells := make([][]int, 0)
	for heightIndex := 0; heightIndex < height; heightIndex++ {
		for widthIndex := 0; widthIndex < width; widthIndex++ {
			if gameMap[heightIndex][widthIndex] == " " {
				position := []int{heightIndex, widthIndex}
				emptyCells = append(emptyCells, position)
			}
		}
	}

	visitedCells := make([][]int, 0)
	currentPosition := make([]int, 2)

	currentPosition[0], currentPosition[1] = startingPosition[0], startingPosition[1]
	// for len(emptyCells) > 0 {
	//REMEMBER TO CHANGE THIS FOR LOOP CONDITION
	for i := 0; i < 100; i++ {
		//TIME LAG FOR VISUAL PURPOSES
		time.Sleep(500 * time.Millisecond)
		gameMap[currentPosition[0]][currentPosition[1]] = " "
		if rand.Intn(2) == 1 {
			//Vertical
			fmt.Println("VERTICAL")
			if rand.Intn(2) == 1 {
				//Going Up
				//If it does not go over the maze
				if currentPosition[0]-2 > 0 {
					fmt.Println("UP")
					//Remove the wall between
					gameMap[currentPosition[0]-1][currentPosition[1]] = " "
					//Variable to add into visitedCells later
					currentPosition[0] = currentPosition[0] - 2
				}
			} else {
				//Going Down
				if currentPosition[0]+2 < height {
					fmt.Println("DOWN")
					//Remove the wall between
					gameMap[currentPosition[0]+1][currentPosition[1]] = " "
					//Variable to add into visitedCells later
					currentPosition[0] = currentPosition[0] + 2
				}
			}
		} else {
			//Horizontal
			//If it does not go over the maze
			fmt.Println("HORIZONTAL")

			if rand.Intn(2) == 1 {
				if currentPosition[1]+2 < width {
					fmt.Println("RIGHT")
					//Remove the wall between
					gameMap[currentPosition[0]][currentPosition[1]+1] = " "
					//Variable to add into visitedCells later
					currentPosition[1] = currentPosition[1] + 2
				}
			} else {
				if currentPosition[1]-2 > 0 {
					fmt.Println("LEFT")
					//Remove the wall between
					gameMap[currentPosition[0]][currentPosition[1]-1] = " "
					//Variable to add into visitedCells later
					currentPosition[1] = currentPosition[1] - 2
				}
			}
		}

		gameMap[currentPosition[0]][currentPosition[1]] = "X"
		visualizeMaze(gameMap)
		visitedCells = append(visitedCells, currentPosition)
	}

	return gameMap
}

func visualizeMaze(maze [][]string) {
	fmt.Println("----------------------")
	//Clear terminal
	fmt.Print("\033[H\033[2J")
	for i := range maze {
		fmt.Println(maze[i])
	}
}
