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

// Notes
// Maybe add save maze function in the future?

// Global var
var direction string = ""
var numSteps int = 0
var width int = 31
var height int = 31

func main() {
	// This program was designed to have and odd width and height
	fmt.Println("Enter your desired width |odd numbers only :( |")
	fmt.Scanln(&width)
	fmt.Println("Enter your desired health |odd numbers only as well : ( |")
	fmt.Scanln(&height)
	if width%2 == 0 {
		width++
	}
	if height%2 == 0 {
		height++
	}
	initializeMap := initialize(width, height)
	generateMaze(initializeMap, width, height)
	// visualizeMaze(generatedMaze, width)
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
	if startingPosition[0]%2 == 0 {
		startingPosition[0]++
	}
	if startingPosition[1]%2 == 0 {
		startingPosition[1]++
	}

	gameMap[startingPosition[0]][startingPosition[1]] = "X"

	startTime := time.Now()
	//Able to choose what type of algo in the future
	gameMap = randomAlgo(gameMap, startingPosition, height, width)

	duration := time.Since(startTime)
	fmt.Printf("Time taken to generate maze:%s", duration)
	// walls := [][]int{}
	// walls = append()
	return gameMap
}

func randomAlgo(gameMap [][]string, startingPosition []int, height, width int) [][]string {
	//Move and remove the wall between each other DONE
	//Randomize whether to go vertical or horizontal DONE
	//Empty cells is " ", so find number of " " and we can know the amount of empty cells and then while there is still empty cells
	//we continue the movement?
	//Need to improve movement so that if it rolls a direction but can't meet its inner condition, then redo another random direction again DONE
	//Add condition to see if the direction being move has been visited b4. DONE
	//A unvisited cell would be a cell with all the walls intact?
	//If it has not moved, it will keep looping but I'm currently not sure how to let it determine that it is impossible to move or it's just really
	//unlucky. just check neighbouring cells around current position to see if they are visited or not dumbass wtf
	visitedCells := make([][]int, 0)
	currentPosition := make([]int, 2)

	currentPosition[0], currentPosition[1] = startingPosition[0], startingPosition[1]
	continueLoop := true
	//REMEMBER TO CHANGE THIS FOR LOOP CONDITION
	for continueLoop {
		//TIME LAG FOR VISUAL PURPOSES
		// time.Sleep(100 * time.Millisecond)
		gameMap[currentPosition[0]][currentPosition[1]] = " "
		var moved bool = false
		//Could get stuck forever if you're really that unlucky but it's suppose to be random tho idk
		for !moved {
			// fmt.Scanf("Hello")
			// fmt.Printf("Current position: %d,%d\n", currentPosition[0], currentPosition[1])

			upCondition := currentPosition[0]-2 < 0 || checkWall(gameMap, []int{currentPosition[0] - 2, currentPosition[1]}) < 4
			downCondition := currentPosition[0]+2 > height || checkWall(gameMap, []int{currentPosition[0] + 2, currentPosition[1]}) < 4
			rightCondition := currentPosition[1]+2 > width || checkWall(gameMap, []int{currentPosition[0], currentPosition[1] + 2}) < 4
			leftCondition := currentPosition[1]-2 < 0 || checkWall(gameMap, []int{currentPosition[0], currentPosition[1] - 2}) < 4
			// fmt.Printf("Up:%t, Down:%t, Right:%t, Left:%t\n", upCondition, downCondition, rightCondition, leftCondition)
			// fmt.Printf("%t\n", upCondition && downCondition && rightCondition && leftCondition)

			if upCondition && downCondition && rightCondition && leftCondition {
				// fmt.Println("HELP STEP BRO IM STUCK")
				// fmt.Printf("%d,%d", currentPosition[0], currentPosition[1])
				gameMap[currentPosition[0]][currentPosition[1]] = " "
				// fmt.Printf("%v", visitedCells)
				//NGL I'm not sure about the code below
				if len(visitedCells) >= 1 {
					currentPosition[0], currentPosition[1] = visitedCells[len(visitedCells)-1][0], visitedCells[len(visitedCells)-1][1]
				} else {
					continueLoop = false
					break
				}
				visitedCells = visitedCells[:len(visitedCells)-1]
				// fmt.Println("GOING BACK")
			}

			if rand.Intn(2) == 1 {
				//Vertical
				// fmt.Println("VERTICAL")
				if rand.Intn(2) == 1 {
					//Going Up
					//If it does not go over the maze AND the cell has not been visited
					if currentPosition[0]-2 > 0 && checkWall(gameMap, []int{currentPosition[0] - 2, currentPosition[1]}) == 4 {
						// fmt.Println("UP")
						gameMap[currentPosition[0]-1][currentPosition[1]] = " "
						currentPosition[0] = currentPosition[0] - 2
						moved = true
					}
				} else {
					//Going Down
					if currentPosition[0]+2 < height && checkWall(gameMap, []int{currentPosition[0] + 2, currentPosition[1]}) == 4 {
						// fmt.Println("DOWN")
						gameMap[currentPosition[0]+1][currentPosition[1]] = " "
						currentPosition[0] = currentPosition[0] + 2
						moved = true
					}
				}
			} else {
				//Horizontal
				//If it does not go over the maze
				// fmt.Println("HORIZONTAL")
				if rand.Intn(2) == 1 {
					if currentPosition[1]+2 < width && checkWall(gameMap, []int{currentPosition[0], currentPosition[1] + 2}) == 4 {
						// fmt.Println("RIGHT")
						gameMap[currentPosition[0]][currentPosition[1]+1] = " "
						currentPosition[1] = currentPosition[1] + 2
						moved = true
					}
				} else {
					if currentPosition[1]-2 > 0 && checkWall(gameMap, []int{currentPosition[0], currentPosition[1] - 2}) == 4 {
						// fmt.Println("LEFT")
						gameMap[currentPosition[0]][currentPosition[1]-1] = " "
						currentPosition[1] = currentPosition[1] - 2
						moved = true
					}
				}
			}

		}

		gameMap[currentPosition[0]][currentPosition[1]] = "X"
		visualizeMaze(gameMap, width)
		copiedPosition := []int{currentPosition[0], currentPosition[1]}
		visitedCells = append(visitedCells, copiedPosition)
	}

	return gameMap
}

func checkWall(gameMap [][]string, cellToCheck []int) int {
	var totalWalls int = 0
	// fmt.Println("CHECKWALL")
	// fmt.Printf("%d,%d\n", cellToCheck[0], cellToCheck[1])
	// fmt.Print("\n")
	//Thanks ChatGPT for the line below, i'm not sure my conditions werent working if the width and height is not the same :/
	if cellToCheck[0]+1 >= len(gameMap) || cellToCheck[1] >= len(gameMap[0]) {
		return 0
	}

	//Check top wall
	if gameMap[cellToCheck[0]+1][cellToCheck[1]] == color.HiBlueString("*") {
		totalWalls = totalWalls + 1
	}
	//Check right wall
	if gameMap[cellToCheck[0]][cellToCheck[1]+1] == color.HiBlueString("*") {
		totalWalls = totalWalls + 1
	}
	//Check bottom wall
	if gameMap[cellToCheck[0]-1][cellToCheck[1]] == color.HiBlueString("*") {
		totalWalls = totalWalls + 1
	}
	//Check left wall
	if gameMap[cellToCheck[0]][cellToCheck[1]-1] == color.HiBlueString("*") {
		totalWalls = totalWalls + 1
	}
	// fmt.Println(totalWalls)
	return totalWalls
}

func visualizeMaze(maze [][]string, width int) {
	//Clear terminal
	fmt.Print("\033[H\033[2J")
	for i := 0; i < 1+width*2; i++ {
		fmt.Print("-")
	}
	fmt.Println("\n")
	fmt.Printf("Current number of algorithm steps taken: %d\n", numSteps)
	numSteps++
	for i := range maze {
		fmt.Println(maze[i])
	}
}
