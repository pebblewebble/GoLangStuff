package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
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
}

type Enemy struct {
	maze            Maze
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

// func (enemy *Enemy) dfs(node [2]int, counter int, visited map[[2]int]int, nonVisited [][2]int, graph map[[2]int][][]int) {
// 	fmt.Println("NEWWWWWWWWWWWWWWWWWWWW")
// 	nonVisited = append(nonVisited, node)
// 	// var notBackTrack bool = true
// 	var previousNode []int = []int{node[0] + 1, node[1] + 1}
// 	for len(nonVisited) != 0 {
// 		if node[0] == enemy.maze.currentPosition[0] && node[1] == enemy.maze.currentPosition[1] {
// 			for key, value := range visited {
// 				if value == 1 {
// 					enemy.currentPosition[0] = key[0]
// 					enemy.currentPosition[1] = key[1]
// 					printVisited(visited)
// 					return
// 				}
// 			}
// 		}
// 		enemy.maze.gameMap[node[0]][node[1]] = " "
// 		node[0], node[1] = nonVisited[0][0], nonVisited[0][1]
// 		if int(math.Abs(float64(node[0]-previousNode[0]))) >= 2 ||
// 			int(math.Abs(float64(node[1]-previousNode[1]))) >= 2 {
// 			fmt.Println("TRUEEEEEEEEEEE")
// 			//Selects most recent visited node
// 			largest := sortVisited(visited, true)
// 			//Becomes the most recent visited node
// 			node[0], node[1] = largest[0], largest[1]
// 			//Get the current node's neighbours
// 			key := [2]int{node[0], node[1]}
// 			currentPositionNeighbour := make([][]int, len(graph[key]))
// 			copy(currentPositionNeighbour, graph[key])
// 			// delete(visited, node)
// 			for {
// 				fmt.Println(backTrack(currentPositionNeighbour, node, visited), node)
// 				if backTrack(currentPositionNeighbour, node, visited) != node {
// 					node[0], node[1] = backTrack(currentPositionNeighbour, node, visited)[0], backTrack(currentPositionNeighbour, node, visited)[1]
// 				} else {
// 					fmt.Println("IM BREAKING")
// 					break
// 				}
// 			}
// 			//Add node into nonvisited
// 			nonVisited = nonVisited[1:]
// 			nonVisited = append([][2]int{node}, nonVisited...)
// 		}
// 		previousNode[0], previousNode[1] = node[0], node[1]

// 		enemy.maze.gameMap[node[0]][node[1]] = color.HiRedString("C")
// 		enemy.maze.visualizeMaze()
// 		var userInput string
// 		fmt.Scan(&userInput)
// 		// time.Sleep(500 * time.Millisecond)
// 		fmt.Println(nonVisited)
// 		if _, exists := visited[[2]int(node)]; !exists {
// 			fmt.Println("I dont exist")
// 			//Doesn't exist? add into visited
// 			visited[[2]int(node)] = counter
// 			counter++
// 			//Get the current node's neighbours
// 			key := [2]int{node[0], node[1]}
// 			currentPositionNeighbour := make([][]int, len(graph[key]))
// 			copy(currentPositionNeighbour, graph[key])
// 			//Remove first element
// 			nonVisited = nonVisited[1:]
// 			//For each of the neighbour, add them back unless they've been visited b4
// 			for i := range currentPositionNeighbour {
// 				temp := [2]int{}
// 				copy(temp[:], currentPositionNeighbour[i])
// 				if _, exists := visited[temp]; !exists {
// 					nonVisited = append([][2]int{temp}, nonVisited...)
// 				}
// 			}
// 		}
// 	}
// }

// func backTrack(currentPositionNeighbour [][]int, node [2]int, visited map[[2]int]int) [2]int {
// 	for i := range currentPositionNeighbour {
// 		temp := [2]int{}
// 		copy(temp[:], currentPositionNeighbour[i])
// 		if _, exists := visited[temp]; exists {
// 			node[0], node[1] = temp[0], temp[1]
// 			return node
// 		}
// 	}
// 	return node
// }

//Stuck in loop
// func (enemy *Enemy) dfs(node [2]int, goal [2]int, counter int, visited map[[2]int]int, nonVisited [][2]int, graph map[[2]int][][]int) {
// 	fmt.Println("NEWWWWWWWWWWWWWWWWWWWW")
// 	fmt.Println(goal)
// 	nonVisited = append(nonVisited, node)
// 	// var notBackTrack bool = true
// 	var previousNode []int = []int{node[0] + 1, node[1] + 1}
// 	for len(nonVisited) != 0 {
// 		fmt.Println("current")
// 		fmt.Println(node)
// 		fmt.Println("goal")
// 		fmt.Println(goal)
// 		if node[0] == goal[0] && node[1] == goal[1] {
// 			for key, value := range visited {
// 				if value == 1 {
// 					enemy.currentPosition[0] = key[0]
// 					enemy.currentPosition[1] = key[1]
// 					printVisited(visited)
// 					return
// 				}
// 			}
// 		}
// 		enemy.maze.gameMap[node[0]][node[1]] = " "
// 		node[0], node[1] = nonVisited[0][0], nonVisited[0][1]
// 		if int(math.Abs(float64(node[0]-previousNode[0]))) >= 2 ||
// 			int(math.Abs(float64(node[1]-previousNode[1]))) >= 2 {
// 			visited2 := make(map[[2]int]int)
// 			nonVisited2 := make([][2]int, 0)
// 			counter2 := 0
// 			enemy.dfs([2]int(previousNode), node, counter2, visited2, nonVisited2, graph)
// 		}
// 		previousNode[0], previousNode[1] = node[0], node[1]

// 		enemy.maze.gameMap[node[0]][node[1]] = color.HiRedString("C")
// 		enemy.maze.visualizeMaze()
// 		// var userInput string
// 		// fmt.Scan(&userInput)
// 		time.Sleep(350 * time.Millisecond)
// 		if _, exists := visited[[2]int(node)]; !exists {
// 			//Doesn't exist? add into visited
// 			visited[[2]int(node)] = counter
// 			counter++
// 			//Get the current node's neighbours
// 			key := [2]int{node[0], node[1]}
// 			currentPositionNeighbour := make([][]int, len(graph[key]))
// 			copy(currentPositionNeighbour, graph[key])
// 			//Remove first element
// 			nonVisited = nonVisited[1:]
// 			//For each of the neighbour, add them back unless they've been visited b4
// 			for i := range currentPositionNeighbour {
// 				temp := [2]int{}
// 				copy(temp[:], currentPositionNeighbour[i])
// 				if _, exists := visited[temp]; !exists {
// 					nonVisited = append([][2]int{temp}, nonVisited...)
// 				}
// 			}
// 		}
// 	}
// }

func (enemy *Enemy) dfs(node [2]int, goal [2]int, counter int, visited [][]int, nonVisited [][2]int, graph map[[2]int][][]int) {
	nonVisited = append(nonVisited, node)
	var previousNode []int = []int{node[0] + 1, node[1] + 1}
	descAmount := 2
	for len(nonVisited) != 0 {
		if node[0] == goal[0] && node[1] == goal[1] {
			for _, i := range visited {
				if i[0] == goal[0] && i[1] == goal[1] {
					enemy.currentPosition[0] = i[0]
					enemy.currentPosition[1] = i[1]
					return
				}
			}
		}
		enemy.maze.gameMap[node[0]][node[1]] = " "

		key := [2]int{node[0], node[1]}
		currentPositionNeighbour := make([][]int, len(graph[key]))
		copy(currentPositionNeighbour, graph[key])
		allNeighbourVisited := true
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
			allNeighbourVisited = false
		} else {
			allNeighbourVisited = true
		}
		if len(visited) != 0 && allNeighbourVisited {
			fmt.Println(node)
			node[0], node[1] = visited[len(visited)-descAmount][0], visited[len(visited)-descAmount][1]
			fmt.Println(visited, descAmount)
			visited = append(visited, []int{node[0], node[1]})
			descAmount += 2
		} else {
			node[0], node[1] = nonVisited[0][0], nonVisited[0][1]
			//reset descAmount after escaping the deadend
			descAmount = 2
		}

		previousNode[0], previousNode[1] = node[0], node[1]
		enemy.maze.gameMap[node[0]][node[1]] = color.HiRedString("C")
		enemy.maze.visualizeMaze()
		// var userInput string
		// fmt.Scan(&userInput)
		time.Sleep(250 * time.Millisecond)

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
}

func (enemy *Enemy) traverseMap(visited map[[2]int]int) {
	keys := sortVisited(visited)
	for _, key := range keys {
		enemy.maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] = " "
		time.Sleep(time.Millisecond * 300)
		enemy.currentPosition[0], enemy.currentPosition[1] = key[0], key[1]
		enemy.maze.gameMap[key[0]][key[1]] = color.HiRedString("C")
		enemy.maze.visualizeMaze()
	}
}

func sortVisited(visited map[[2]int]int) [][2]int {
	// printVisited(visited)
	// fmt.Println(i)
	keys := make([][2]int, 0, len(visited))

	// Populate the slice with keys from the map
	for key := range visited {
		keys = append(keys, key)
	}

	// Sort the keys based on their corresponding values
	sort.Slice(keys, func(i, j int) bool {
		return visited[keys[i]] < visited[keys[j]]
	})

	return keys

}

func getLargest(visited map[[2]int]int, largest bool, i int) [2]int {
	keys := sortVisited(visited)
	// Return the desired key
	if largest {
		return keys[len(keys)-i]
	} else {
		return keys[0]
	}
}

func printVisited(visited map[[2]int]int) {
	// Create a slice to store the keys for sorting
	keys := sortVisited(visited)

	// Print the keys in ascending order of the counter
	for _, key := range keys {
		fmt.Printf("%v: %d\n", key, visited[key])
	}
}

func (enemy *Enemy) makeGraph() map[[2]int][][]int {
	graph := make(map[[2]int][][]int)
	for i := range enemy.maze.gameMap {
		for x := range enemy.maze.gameMap[0] {
			neighbours := make([][]int, 0)
			if enemy.maze.gameMap[i][x] != color.HiBlueString("*") {
				if enemy.maze.gameMap[i-1][x] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i - 1, x})
				}
				if enemy.maze.gameMap[i+1][x] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i + 1, x})
				}
				if enemy.maze.gameMap[i][x+1] != color.HiBlueString("*") {
					neighbours = append(neighbours, []int{i, x + 1})
				}
				if enemy.maze.gameMap[i][x-1] != color.HiBlueString("*") {
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
		maze.algoSteps++
		// maze.visualizeMaze()
		copiedPosition := []int{maze.currentPosition[0], maze.currentPosition[1]}
		visitedCells = append(visitedCells, copiedPosition)
	}
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = " "
	maze.currentPosition[0], maze.currentPosition[1] = maze.startingPosition[0], maze.startingPosition[1]
	maze.gameMap[maze.currentPosition[0]][maze.currentPosition[1]] = "X"
	return maze.gameMap
}

func (maze *Maze) movement(energyPos [][]int, enemy *Enemy) {
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
		maze.checkExit()
		checkEnergy(energyPos)
		maze.numSteps++
		fmt.Println("Hello")

		visited := [][]int{}
		nonVisited := make([][2]int, 0)
		graph := enemy.makeGraph()
		node := [2]int{enemy.currentPosition[0], enemy.currentPosition[1]}
		goal := [2]int{enemy.maze.currentPosition[0], enemy.maze.currentPosition[1]}
		maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] = " "
		//Updates enemy currentPosition here
		fmt.Println("BEFORE")
		fmt.Println(node)
		enemy.dfs(node, goal, 0, visited, nonVisited, graph)
		fmt.Println("AFTER")
		fmt.Println(enemy.currentPosition)
		//Replace and draw
		maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] = color.HiRedString("C")
		maze.visualizeMaze()
	}
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

func (enemy *Enemy) spawn() {
	for {
		enemy.currentPosition = []int{rand.Intn(len(enemy.maze.gameMap)), rand.Intn(len(enemy.maze.gameMap[0]))}
		if enemy.maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] == " " {
			enemy.maze.gameMap[enemy.currentPosition[0]][enemy.currentPosition[1]] = color.HiRedString("C")
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
	enemy1.maze = *newMaze
	enemy1.spawn()
	energyPos := newMaze.generateEnergy(numOrb)
	continueMovement := false
	for !continueMovement {
		newMaze.movement(energyPos, enemy1)
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
	// fmt.Print("\033[H\033[2J")
	// fmt.Printf("\033[%d;%dH", 0+1, 0+1)
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
