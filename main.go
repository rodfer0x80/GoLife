package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Configurations
const TOTAL_CELLS int = 16
const HIVE_SIZE int = TOTAL_CELLS + 1

// ----

func genesis() string {
	rand.Seed(time.Now().UnixNano())
	var n int = rand.Intn(2)

	var char string = ""
	if n == 1 {
		char = "x"
	} else {
		char = "o"
	}

	return char
}

func bigBang(hive *[HIVE_SIZE]string, generation int) (*[HIVE_SIZE]string, int) {
	generation = 1
	for i := 1; i <= TOTAL_CELLS; i++ {
		hive[i] = genesis()
	}

	return hive, generation
}

func tick(cell string, hive *[HIVE_SIZE]string) string {
	if cell == "x" {
		cell = "o"
	} else {
		cell = "x"
	}
	return cell
}

func naturalSelection(hive *[HIVE_SIZE]string, generation int) (*[HIVE_SIZE]string, int) {
	var postHive = new([HIVE_SIZE]string)

	generation = generation + 1
	for i := 1; i <= TOTAL_CELLS; i++ {
		postHive[i] = tick(hive[i], hive)
	}

	for i := 1; i <= TOTAL_CELLS; i++ {
		hive[i] = postHive[i]
	}

	return hive, generation
}

func displayGrid(hive *[HIVE_SIZE]string, generation int) int {
	var rowCells int = int(math.Sqrt(float64(TOTAL_CELLS)))
	var columnCells int = rowCells

	fmt.Println("Generation:" + strconv.Itoa(generation))
	for i := 1; i <= columnCells; i++ {
		buffer := ""
		for ii := 1; ii <= rowCells; ii++ {
			buffer += hive[i*ii] + " "
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
