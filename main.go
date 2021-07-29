package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Graphical configurations
const alive string = "o"
const dead string = "."

// const alive string = "●"
// const dead string = "○"

// ----

func clearScreen() int {
	fmt.Println("\033[2J")
	return 0
}

func genesis() string {
	rand.Seed(time.Now().UnixNano())
	var randomNumber int = rand.Intn(2)

	var state string = ""
	if randomNumber == 1 {
		state = dead
	} else {
		state = alive
	}

	return state
}

func bigBang(hive []string, generation int, TOTAL_CELLS int) int {
	generation = 1
	for i := 0; i < TOTAL_CELLS; i++ {
		hive[i] = genesis()
	}

	return generation
}

func naturalSelection(cell string, neighbours int) string {
	if (cell == alive) && ((neighbours < 2) || (neighbours > 3)) {
		cell = dead
	}
	if (cell == dead) && (neighbours == 3) {
		cell = alive
	}

	return cell
}

func outOfBoardRight(column int, rowCells int) bool {
	var flag bool = false

	if column == rowCells-1 {
		flag = true
	}

	return flag
}

func outOfBoardLeft(column int) bool {
	var flag bool = false

	if column == 0 {
		flag = true
	}

	return flag
}

func outOfBoardUp(row int) bool {
	var flag bool = false

	if row == 0 {
		flag = true
	}

	return flag
}

func outOfBoardDown(row int, rowCells int) bool {
	var flag bool = false

	if row == rowCells-1 {
		flag = true
	}

	return flag
}

func getNeighbourhood(row int, column int, rowCells int, position int, hive []string) int {
	var neighbours int = 0

	// Left
	if !outOfBoardLeft(column) {
		if hive[position-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Right
	if !outOfBoardRight(column, rowCells) {
		if hive[position+1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Up
	if !outOfBoardUp(row) {
		if hive[position-rowCells] == alive {
			neighbours = neighbours + 1
		}
	}

	// Down
	if !outOfBoardDown(row, rowCells) {
		if hive[position+rowCells] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal left up
	if (!outOfBoardLeft(column)) && (!outOfBoardUp(row)) {
		if hive[position-rowCells-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal right up
	if (!outOfBoardRight(column, rowCells)) && (!outOfBoardUp(row)) {
		if hive[position-rowCells+1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal left down
	if (!outOfBoardLeft(column)) && (!outOfBoardDown(row, rowCells)) {
		if hive[position+rowCells-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal right down
	if (!outOfBoardRight(column, rowCells)) && (!outOfBoardDown(row, rowCells)) {
		if hive[position+rowCells+1] == alive {

			neighbours = neighbours + 1
		}
	}

	return neighbours
}

func tick(hive []string, position int, TOTAL_CELLS int) string {
	var rowCells int = int(math.Sqrt(float64(TOTAL_CELLS)))
	var row int = position / rowCells
	var column int = position - (rowCells * row)

	var neighbours int = getNeighbourhood(row, column, rowCells, position, hive)

	var cell string = hive[position]
	cell = naturalSelection(cell, neighbours)

	return cell
}

func cycleOfLife(hive []string, generation int, TOTAL_CELLS int) int {
	var postHive = make([]string, TOTAL_CELLS, TOTAL_CELLS)

	generation = generation + 1

	for position := 0; position < TOTAL_CELLS; position++ {
		postHive[position] = tick(hive, position, TOTAL_CELLS)
	}

	for position := 0; position < TOTAL_CELLS; position++ {
		hive[position] = postHive[position]
	}

	return generation
}

func displayGrid(hive []string, generation int, TOTAL_CELLS int) int {
	var rowCells int = int(math.Sqrt(float64(TOTAL_CELLS)))

	var buffer string

	clearScreen()
	fmt.Println("Generation:" + strconv.Itoa(generation))
	for i := 0; i < TOTAL_CELLS; i = i + rowCells {
		buffer = ""
		for ii := 0; ii < rowCells; ii++ {
			buffer = buffer + hive[i+ii] + " "
		}
		fmt.Println(buffer)
	}

	return 0
}

func main() {
	// should be a perfect square
	const TOTAL_CELLS int = 900

	var generation int = 0
	var hive = make([]string, TOTAL_CELLS, TOTAL_CELLS)

	generation = bigBang(hive[:], generation, TOTAL_CELLS)
	displayGrid(hive[:], generation, TOTAL_CELLS)

	for 1 < 2 {
		time.Sleep(1 * time.Second)
		generation = cycleOfLife(hive[:], generation, TOTAL_CELLS)
		displayGrid(hive[:], generation, TOTAL_CELLS)
	}
}
