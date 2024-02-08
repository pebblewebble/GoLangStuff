package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
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
// var numSteps int = 0
// var width int = 31
// var height int = 31

type Maze struct {
	width            int
	height           int
	numSteps         int
	startingPosition []int
	currentPosition  []int
	gameMap          [][]string
	exit             []int
	algoSteps        int
	enemy            []*Enemy
}

type Enemy struct {
	currentPosition []int
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

	maze.gameMap[maze.startingPosition[0]][maze.startingPosition[1]] = "X"

	startTime := time.Now()
	//Able to choose what type of algo in the future
	maze.gameMap = maze.randomAlgo()

	duration := time.Since(startTime)
	fmt.Printf("Time taken to generate maze:%s", duration)
	return maze.gameMap
}

func (enemy *Enemy) greedyBestFirst(visited, nonVisited map[int][]int) {
	// //Neighbours
	// topCellHeuristic := getTotalDistance([]int{enemy.currentPosition[0] - 1, enemy.currentPosition[1]}, enemy.maze.currentPosition)

}

func getTotalDistance(enemyPos, playerPos []int) int {
	heightDistance := math.Abs(math.Inf(enemyPos[0] - playerPos[0]))
	widthDistance := math.Abs(math.Inf(enemyPos[1] - playerPos[1]))
	return int(heightDistance + widthDistance)
}

func (enemy *Enemy) dfs(node [2]int, goal [2]int, counter int, visited [][]int, nonVisited [][2]int, graph map[[2]int][][]int) string {
	nonVisited = append(nonVisited, node)
	callStack := [][]int{}
	descAmount := 2
	for len(nonVisited) != 0 {
		if node[0] == goal[0] && node[1] == goal[1] {
			enemy.currentPosition[0] = callStack[1][0]
			enemy.currentPosition[1] = callStack[1][1]
			if enemy.currentPosition[0] == goal[0] && enemy.currentPosition[1] == goal[1] {
				return "LOSE"
			}
			return ""
		}

		allNeighbourVisited := checkStuck(node, graph, visited)

		if len(visited) != 0 && allNeighbourVisited && descAmount <= len(visited) {
			node[0], node[1] = visited[len(visited)-descAmount][0], visited[len(visited)-descAmount][1]
			//This part sometimes goes out of bounds
			//I think the issue is that, for example the entire left side is stuck, and the only way to go to
			//the goal is through the original spawn point which will cause the callStack to have 0?
			//This method sometimes teleports idk why - 6/2/2024
			//Nvm there is still some problems
			if len(callStack) != 1 {
				callStack = callStack[:len(callStack)-1]
			}
			descAmount += 1
		} else {
			node[0], node[1] = nonVisited[0][0], nonVisited[0][1]
			//reset descAmount after escaping the deadend
			descAmount = 1
		}

		existsCheck := false
		for _, i := range visited {
			if i[0] == node[0] && i[1] == node[1] {
				existsCheck = true
				break
			}
		}
		if !existsCheck {
			//Doesn't exist? add into visited
			visited = append(visited, []int{node[0], node[1]})
			callStack = append(callStack, []int{node[0], node[1]})
			//Get the current node's neighbours
			key := [2]int{node[0], node[1]}
			currentPositionNeighbour := make([][]int, len(graph[key]))
			copy(currentPositionNeighbour, graph[key])
			//Remove first element
			nonVisited = nonVisited[1:]
			//For each of the neighbour, add them back unless they've been visited b4
			for i := range currentPositionNeighbour {
				temp := [2]int{}
				copy(temp[:], currentPositionNeighbour[i])
				existsCheck = false
				for _, i := range visited {
					if i[0] == temp[0] && i[1] == temp[1] {
						existsCheck = true
						break
					}
				}
				if !existsCheck {
					nonVisited = append([][2]int{temp}, nonVisited...)
				}
			}
		}
	}
	return ""
}

func checkStuck(node [2]int, graph map[[2]int][][]int, visited [][]int) bool {
	key := [2]int{node[0], node[1]}
	currentPositionNeighbour := make([][]int, len(graph[key]))
	copy(currentPositionNeighbour, graph[key])
	counter := len(currentPositionNeighbour)
	for i := range currentPositionNeighbour {
		if counter == 0 {
			break
		}
		temp := [2]int{}
		copy(temp[:], currentPositionNeighbour[i])
		//Checks all of visited against all of current neighbour
		for _, x := range visited {
			if x[0] == temp[0] && x[1] == temp[1] {
				counter -= 1
				break
			}
		}
	}
	if counter != 0 {
		return false
	} else {
		return true
	}
}

func (maze *Maze) makeGraph() map[[2]int][][]int {
	graph := make(map[[2]int][][]int)
	for i := range maze.gameMap {
		for x := range maze.gameMap[0] {
			neighbours := make([][]int, 0)
			if maze.gameMap[i][x] != color.HiBlueString("*") {
				if maze.gameMap[i-1][x] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i - 1, x})
				}
				if maze.gameMap[i+1][x] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i + 1, x})
				}
				if maze.gameMap[i][x+1] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i, x + 1})
				}
				if maze.gameMap[i][x-1] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i, x - 1})
				}
				graph[[2]int{i, x}] = neighbours
			}
		}
	}
	return graph
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
		// time.Sleep(100 * time.Millisecond)
		maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
		var moved bool = false
		//Could get stuck forever if you're really that unlucky but it's suppose to be random tho idk
		for !moved {
			upCondition := maze.currentPosition[0]-2 < 0 || checkWall(maze.gameMap, []int{maze.currentPosition[0] - 2, maze.currentPosition[1]}) < 4
			downCondition := maze.currentPosition[0]+2 > maze.height || checkWall(maze.gameMap, []int{maze.currentPosition[0] + 2, maze.currentPosition[1]}) < 4
			rightCondition := maze.currentPosition[1]+2 > maze.width || checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] + 2}) < 4
			leftCondition := maze.currentPosition[1]-2 < 0 || checkWall(maze.gameMap, []int{maze.currentPosition[0], maze.currentPosition[1] - 2}) < 4

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
		maze.algoSteps++
		copiedPosition := []int{maze.currentPosition[0], maze.currentPosition[1]}
		visitedCells = append(visitedCells, copiedPosition)
	}
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
	maze.currentPosition[0], maze.currentPosition[1] = maze.startingPosition[0], maze.startingPosition[1]
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = "X"
	return maze.gameMap
}

func (maze *Maze) movement(energyPos [][]int, enemy []*Enemy) string {
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
		state := ""
		switch string(r) {
		case "w":
			if maze.checkMovement("w") {
				maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[0] = maze.currentPosition[0] - 1
			}
			state = maze.afterMovementLogic(energyPos, enemy)
		case "a":
			if maze.checkMovement("a") {
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]-1] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[1] = maze.currentPosition[1] - 1
			}
			state = maze.afterMovementLogic(energyPos, enemy)
		case "s":
			if maze.checkMovement("s") {
				maze.gameMap[maze.currentPosition[0]+1][maze.currentPosition[1]] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[0] = maze.currentPosition[0] + 1
			}
			state = maze.afterMovementLogic(energyPos, enemy)
		case "d":
			if maze.checkMovement("d") {
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]+1] = "X"
				maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
				maze.currentPosition[1] = maze.currentPosition[1] + 1
			}
			state = maze.afterMovementLogic(energyPos, enemy)
		case "q":
			return "QUIT"
		}
		if state == "LOSE" {
			return "LOSE"
		}
	}
}

func (maze *Maze) afterMovementLogic(energyPos [][]int, enemy []*Enemy) string {
	//BTW this was made because if it is after the switch case, it will compute twice, meaning like the step
	// will +2 instead.
	maze.numSteps++
	maze.checkExit()

	checkEnergy(energyPos)

	visited := [][]int{}
	nonVisited := make([][2]int, 0)
	graph := maze.makeGraph()
	goalCheck := ""
	for _, value := range enemy {
		node := [2]int{value.currentPosition[0], value.currentPosition[1]}
		goal := [2]int{maze.currentPosition[0], maze.currentPosition[1]}
		maze.gameMap[value.currentPosition[0]][value.currentPosition[1]] = " "
		//Updates enemy currentPosition here
		goalCheck = value.dfs(node, goal, 0, visited, nonVisited, graph)
		if goalCheck == "LOSE" {
			return "LOSE"
		}
		//Replace and draw
		maze.gameMap[value.currentPosition[0]][value.currentPosition[1]] = color.HiRedString("C")
	}

	maze.visualizeMaze()
	return "CONTINUE"
}

func gameOver() {
	fmt.Println("YOU LOST")
}

func (maze *Maze) generateStart() []int {
	maze.startingPosition = []int{rand.Intn(maze.height-1-1) + 1, rand.Intn(maze.width-2) + 1}
	return maze.startingPosition
}

func (maze *Maze) generateExit() {
	for {
		maze.exit = []int{rand.Intn(len(maze.gameMap)), rand.Intn(len(maze.gameMap[0]))}
		if maze.gameMap[maze.exit[0]][maze.exit[1]] != color.HiBlueString("*") {
			maze.gameMap[maze.exit[0]][maze.exit[1]] = color.HiYellowString("O")
			break
		}
	}
}

func (maze *Maze) spawn(enemy *Enemy) {
	for {
		enemy.currentPosition = []int{rand.Intn(len(maze.gameMap)), rand.Intn(len(maze.gameMap[0]))}
		if maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] == " " {
			maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] = color.HiRedString("C")
			break
		}
	}
}

func (maze *Maze) generateEnergy(numberOfOrbs int) [][]int {
	energyPos := make([][]int, numberOfOrbs)
	for i := 0; i < numberOfOrbs; i++ {
		for {
			mazeEnergy := []int{rand.Intn(len(maze.gameMap)), rand.Intn(len(maze.gameMap[0]))}
			if maze.gameMap[mazeEnergy[0]][mazeEnergy[1]] != color.HiBlueString("*") &&
				maze.gameMap[mazeEnergy[0]][mazeEnergy[1]] != "X" &&
				maze.gameMap[mazeEnergy[0]][mazeEnergy[1]] != color.HiBlueString("O") {
				maze.gameMap[mazeEnergy[0]][mazeEnergy[1]] = color.HiGreenString("E")
				energyPos = append(energyPos, mazeEnergy)
				break
			}
		}
	}
	return energyPos
}

func (maze *Maze) checkExit() {
	if maze.gameMap[maze.exit[0]][maze.exit[1]] != color.HiYellowString("O") {
		maze.generateExit()
	}
}

func checkEnergy(energyCords [][]int) {

}

func main() {
	// This program was designed to have and odd width and height
	for {
		fmt.Println("Enter your desired width ")
		newMaze := new(Maze)
		fmt.Scanln(&newMaze.width)
		fmt.Println("Enter your desired height ")
		fmt.Scanln(&newMaze.height)
		fmt.Println("Enter your desired energy orbs ")
		numOrb := 0
		fmt.Scanln(&numOrb)
		if newMaze.width%2 == 0 {
			newMaze.width++
		}
		if newMaze.height%2 == 0 {
			newMaze.height++
		}
		newMaze.initialize()
		newMaze.generateMaze()
		newMaze.generateExit()
		enemy1 := new(Enemy)
		enemy2 := new(Enemy)
		newMaze.enemy = append(newMaze.enemy, enemy1, enemy2)
		for i := range newMaze.enemy {
			newMaze.spawn(newMaze.enemy[i])
		}
		energyPos := newMaze.generateEnergy(numOrb)
		newMaze.visualizeMaze()
		continueMovement := false
		stateCheck := ""
		for !continueMovement {
			stateCheck = newMaze.movement(energyPos, newMaze.enemy)
			if stateCheck == "LOSE" {
				break
			}
		}
	}
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
		return maze.currentPosition[0]-1 > 0 && maze.gameMap[maze.currentPosition[0]-1][maze.currentPosition[1]] != color.HiBlueString("*")
	case "a":
		// Check if moving left is valid
		return maze.currentPosition[1]-1 > 0 && maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]-1] != color.HiBlueString("*")
	case "s":
		// Check if moving down is valid
		return maze.currentPosition[0]+1 < maze.height && maze.gameMap[maze.currentPosition[0]+1][maze.currentPosition[1]] != color.HiBlueString("*")
	case "d":
		// Check if moving right is valid
		return maze.currentPosition[1]+1 < maze.width && maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]+1] != color.HiBlueString("*")
	default:
		// Invalid movement key
		return false
	}
}

func (maze *Maze) visualizeMaze() {
	//Clear terminal
	fmt.Print("\033[H\033[2J")
	fmt.Printf("\033[%d;%dH", 0+1, 0+1)
	for i := 0; i < 1+maze.width*2; i++ {
		fmt.Print("-")
	}
	fmt.Println("\n")
	fmt.Printf("Number of steps to create maze: %d\n", maze.algoSteps)
	fmt.Printf("Number of steps you have taken: %d\n", maze.numSteps)
	for i := range maze.gameMap {
		fmt.Println(maze.gameMap[i])
	}
}
