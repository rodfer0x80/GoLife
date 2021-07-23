package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Configurations
const TOTAL_CELLS int = 1600
const alive string = "O"
const dead string = "."

// ----

// Code readibility
const HIVE_SIZE int = TOTAL_CELLS + 1

// ----

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

func bigBang(hive *[HIVE_SIZE]string, generation int) (*[HIVE_SIZE]string, int) {
	generation = 1
	for i := 1; i <= TOTAL_CELLS; i++ {
		hive[i] = genesis()
	}

	return hive, generation
}

func cycleOfLife(cell string, neighbours int) string {
	if (cell == alive) && ((neighbours < 2) || (neighbours > 3)) {
		cell = dead
	}
	if (cell == dead) && (neighbours == 3) {
		cell = alive
	}

	return cell
}
func tick(hive *[HIVE_SIZE]string, position int) string {
	var rowCells int = int(math.Sqrt(float64(TOTAL_CELLS)))
	var neighbours int = 0

	var outOfBoardLeft bool
	var outOfBoardRight bool
	var outOfBoardUp bool
	var outOfBoardDown bool
	var boardPosition int

	// Left
	outOfBoardLeft = false
	boardPosition = 1
	for boardPosition < HIVE_SIZE {
		if position == boardPosition {
			outOfBoardLeft = true
		}
		boardPosition = boardPosition + rowCells
	}
	if outOfBoardLeft == false {
		if hive[position-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Right
	outOfBoardRight = false
	boardPosition = rowCells
	for boardPosition < HIVE_SIZE {
		if position == boardPosition {
			outOfBoardRight = true
		}
		boardPosition = boardPosition + rowCells
	}
	if outOfBoardRight == false {
		if hive[position+1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Up
	outOfBoardUp = false
	boardPosition = 1
	for boardPosition <= rowCells {
		if position == boardPosition {
			outOfBoardUp = true
		}
		boardPosition = boardPosition + 1
	}
	if outOfBoardUp == false {
		if hive[position-rowCells] == alive {
			neighbours = neighbours + 1
		}
	}

	// Down
	outOfBoardDown = false
	boardPosition = (HIVE_SIZE - rowCells)
	for boardPosition < HIVE_SIZE {
		if position == boardPosition {
			outOfBoardDown = true
		}
		boardPosition = boardPosition + 1
	}
	if outOfBoardDown == false {
		if hive[position+rowCells] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal left up
	if (outOfBoardLeft == false) && (outOfBoardUp == false) {
		if hive[position-rowCells-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal right up
	if (outOfBoardRight == false) && (outOfBoardUp == false) {
		if hive[position-rowCells+1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal left down
	if (outOfBoardLeft == false) && (outOfBoardDown == false) {
		if hive[position+rowCells-1] == alive {
			neighbours = neighbours + 1
		}
	}

	// Diagonal right now
	if (outOfBoardRight == false) && (outOfBoardDown == false) {
		if hive[position+rowCells+1] == alive {
			neighbours = neighbours + 1
		}
	}

	var cell string = hive[position]
	cell = cycleOfLife(cell, neighbours)

	return cell
}

func naturalSelection(hive *[HIVE_SIZE]string, generation int) (*[HIVE_SIZE]string, int) {
	var postHive = new([HIVE_SIZE]string)

	generation = generation + 1
	for position := 1; position <= TOTAL_CELLS; position++ {
		postHive[position] = tick(hive, position)
	}

	for position := 1; position <= TOTAL_CELLS; position++ {
		hive[position] = postHive[position]
	}

	return hive, generation
}

func displayGrid(hive *[HIVE_SIZE]string, generation int) int {
	var rowCells int = int(math.Sqrt(float64(TOTAL_CELLS)))
	var buffer string

	fmt.Println("Generation:" + strconv.Itoa(generation))
	for i := 0; i < TOTAL_CELLS; i = i + rowCells {
		buffer = ""
		for ii := 1; ii <= rowCells; ii++ {
			buffer = buffer + hive[i+ii] + " "
		}
		fmt.Println(buffer)
	}

	return 0
}

func main() {
	var generation int = 0
	var hive = new([HIVE_SIZE]string)

	hive, generation = bigBang(hive, generation)
	displayGrid(hive, generation)

	for 1 < 2 {
		time.Sleep(1 * time.Second)
		hive, generation = naturalSelection(hive, generation)
		displayGrid(hive, generation)
	}
}
