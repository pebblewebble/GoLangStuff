package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/mattn/go-tty"
)

// REFERENCES
// https://www.baeldung.com/cs/maze-generation#:~:text=Another%20way%20of%20generating%20a,as%20visited%2C%20and%20so%20on.
// https://medium.com/swlh/fun-with-python-1-maze-generator-931639b4fb7e

// Notes
// Maybe add save maze function in the future?

// Global var
var numSteps int = 0

// var width int = 31
// var height int = 31

type Maze struct {
	width            int
	height           int
	numSteps         int
	startingPosition []int
	currentPosition  []int
	gameMap          [][]string
}

func (maze *Maze) initialize() [][]string {
	//Creating the matrix
	maze.gameMap = make([][]string, maze.height)
	for heightIndex := 0; heightIndex < maze.height; heightIndex++ {
		maze.gameMap[heightIndex] = make([]string, maze.width)
		for widthIndex := 0; widthIndex < maze.width; widthIndex++ {
			//Generate outer walls and cells
			if heightIndex == 0 || heightIndex == maze.height-1 || widthIndex == 0 || widthIndex == maze.width-1 || heightIndex%2 == 0 || widthIndex%2 == 0 {
				maze.gameMap[heightIndex][widthIndex] = color.HiBlueString("*")
			} else {
				maze.gameMap[heightIndex][widthIndex] = " "
			}
		}
	}

	return maze.gameMap
}

func (maze *Maze) generateMaze() [][]string {
	//Generates a random odd X,Y starting position but instead of constantly randomly generating a random position each time,
	//I could just minus 1 from the position that is not odd.
	maze.generateStart()
	if maze.startingPosition[0]%2 == 0 {
		maze.startingPosition[0]++
	}
	if maze.startingPosition[1]%2 == 0 {
		maze.startingPosition[1]++
	}

	fmt.Println(maze.gameMap)
	maze.gameMap[maze.startingPosition[0]][maze.startingPosition[1]] = "X"

	startTime := time.Now()
	//Able to choose what type of algo in the future
	maze.gameMap = maze.randomAlgo()

	duration := time.Since(startTime)
	fmt.Printf("Time taken to generate maze:%s", duration)
	return maze.gameMap
}

func (maze *Maze) randomAlgo() [][]string {
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
	maze.currentPosition = make([]int, 2)

	maze.currentPosition[0], maze.currentPosition[1] = maze.startingPosition[0], maze.startingPosition[1]
	continueLoop := true
	//REMEMBER TO CHANGE THIS FOR LOOP CONDITION
	for continueLoop {
		//TIME LAG FOR VISUAL PURPOSES
		// time.Sleep(500 * time.Millisecond)
		maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
		var moved bool = false
		//Could get stuck forever if you're really that unlucky but it's suppose to be random tho idk
		for !moved {
			// fmt.Scanf("Hello")
			// fmt.Printf("Current position: %d,%d\n", currentPosition[0], currentPosition[1])

			upCondition := maze.currentPosition[0]-2 < 0 || checkWall(maze.gameMap, []int{maze.currentPosition[0] - 2, maze.currentPosition[1]}) < 4
			downCondition := maze.currentPosition[0]+2 > maze.height || checkWall(maze.gameMap, []int{maze.currentPosition[0] + 2, maze.currentPosition[1]}) < 4
			rightCondition := maze.currentPosition[1]+2 > maze.width || checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] + 2}) < 4
			leftCondition := maze.currentPosition[1]-2 < 0 || checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] - 2}) < 4
			// fmt.Printf("Up:%t, Down:%t, Right:%t, Left:%t\n", upCondition, downCondition, rightCondition, leftCondition)
			// fmt.Printf("%t\n", upCondition && downCondition && rightCondition && leftCondition)

			if upCondition && downCondition && rightCondition && leftCondition {
				// fmt.Println("HELP STEP BRO IM STUCK")
				// fmt.Printf("%d,%d", currentPosition[0], currentPosition[1])
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				// fmt.Printf("%v", visitedCells)
				//NGL I'm not sure about the code below
				if len(visitedCells) >= 1 {
					maze.currentPosition[0], maze.currentPosition[1] = visitedCells[len(visitedCells)-1][0], visitedCells[len(visitedCells)-1][1]
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
					if maze.currentPosition[0]-2 > 0 && checkWall(maze.gameMap, []int{maze.currentPosition[0] - 2, maze.currentPosition[1]}) == 4 {
						// fmt.Println("UP")
						maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] = " "
						maze.currentPosition[0] = maze.currentPosition[0] - 2
						moved = true
					}
				} else {
					//Going Down
					if maze.currentPosition[0]+2 < maze.height && checkWall(maze.gameMap, []int{maze.currentPosition[0] + 2, maze.currentPosition[1]}) == 4 {
						// fmt.Println("DOWN")
						maze.gameMap[maze.currentPosition[0]+1][maze.currentPosition[1]] = " "
						maze.currentPosition[0] = maze.currentPosition[0] + 2
						moved = true
					}
				}
			} else {
				//Horizontal
				//If it does not go over the maze
				// fmt.Println("HORIZONTAL")
				if rand.Intn(2) == 1 {
					if maze.currentPosition[1]+2 < maze.width && checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] + 2}) == 4 {
						// fmt.Println("RIGHT")
						maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]+1] = " "
						maze.currentPosition[1] = maze.currentPosition[1] + 2
						moved = true
					}
				} else {
					if maze.currentPosition[1]-2 > 0 && checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] - 2}) == 4 {
						// fmt.Println("LEFT")
						maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]-1] = " "
						maze.currentPosition[1] = maze.currentPosition[1] - 2
						moved = true
					}
				}
			}

		}

		maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = "X"
		maze.visualizeMaze()
		copiedPosition := []int{maze.currentPosition[0], maze.currentPosition[1]}
		visitedCells = append(visitedCells, copiedPosition)
	}
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
	maze.currentPosition[0], maze.currentPosition[1] = maze.startingPosition[0], maze.startingPosition[1]
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = "X"
	return maze.gameMap
}

func (maze *Maze) movement() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch string(r) {
		case "w":
			if maze.checkMovement("w") {
				maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[0] = maze.currentPosition[0] - 1
			}
		case "a":
			if maze.checkMovement("a") {
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]-1] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[1] = maze.currentPosition[1] - 1
			}
		case "s":
			if maze.checkMovement("s") {
				maze.gameMap[maze.currentPosition[0]+1][maze.currentPosition[1]] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[0] = maze.currentPosition[0] + 1
			}
		case "d":
			if maze.checkMovement("d") {
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]+1] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[1] = maze.currentPosition[1] + 1
			}
		case "q":
			os.Exit(3)
		}
		maze.visualizeMaze()
		//Displays which key the user pressed
		// fmt.Println("\nKey press => " + string(r))
	}
}

func (maze *Maze) generateStart() []int {
	maze.startingPosition = []int{rand.Intn(maze.height-1-1) + 1, rand.Intn(maze.width-2) + 1}
	return maze.startingPosition
}

func main() {
	// This program was designed to have and odd width and height
	fmt.Println("Enter your desired width ")
	newMaze := Maze{numSteps: 0}
	fmt.Scanln(&newMaze.width)
	fmt.Println("Enter your desired health ")
	fmt.Scanln(&newMaze.height)
	if newMaze.width%2 == 0 {
		newMaze.width++
	}
	if newMaze.height%2 == 0 {
		newMaze.height++
	}
	newMaze.initialize()
	newMaze.generateMaze()
	continueMovement := false
	for !continueMovement {
		newMaze.movement()
	}
	// visualizeMaze(generatedMaze, width)
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

func (maze Maze) checkMovement(movementKey string) bool {
	// if movementKey == "w" && maze.currentPosition[0]-1 > 0 && maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] == " " {
	// 	return true
	// }
	// return false
	//ChatGPT optimised my method below, original on top
	switch movementKey {
	case "w":
		// Check if moving up is valid
		return maze.currentPosition[0]-1 > 0 && maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] == " "
	case "a":
		// Check if moving left is valid
		return maze.currentPosition[1]-1 > 0 && maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]-1] == " "
	case "s":
		// Check if moving down is valid
		return maze.currentPosition[0]+1 < maze.height && maze.gameMap[maze.currentPosition[0]+1][maze.currentPosition[1]] == " "
	case "d":
		// Check if moving right is valid
		return maze.currentPosition[1]+1 < maze.width && maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]+1] == " "
	default:
		// Invalid movement key
		return false
	}
}

func (maze *Maze) visualizeMaze() {
	//Clear terminal
	// fmt.Print("\033[H\033[2J")
	for i := 0; i < 1+maze.width*2; i++ {
		fmt.Print("-")
	}
	fmt.Println("\n")
	fmt.Printf("Current number of algorithm steps taken: %d\n", maze.numSteps)
	maze.numSteps++
	for i := range maze.gameMap {
		fmt.Println(maze.gameMap[i])
	}
}
