package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

//func Tick([]int) int {
//	checkAlive(cell bool)
//}
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

func bigBang(postState *[17]string, generation int, totalCells int) (*[17]string, int) {
	generation = 1
	for i := 1; i <= totalCells; i++ {
		postState[i] = genesis()
	}

	return postState, generation
}

func isAlive(cell string) string {
	if cell == "x" {
		cell = "o"
	} else {
		cell = "x"
	}
	return cell
}

func displayGrid(postState *[17]string, generation int, totalCells int, rowCells int, columnCells int) int {
	fmt.Println("Generation:" + strconv.Itoa(generation))
	for i := 1; i <= columnCells; i++ {
		buffer := ""
		for ii := 1; ii <= rowCells; ii++ {
			buffer += postState[i*ii] + " "
		}
		fmt.Println(buffer)
	}

	return 0
}

func main() {
	const totalCells int = 16
	var rowCells int = int(math.Sqrt(float64(totalCells)))
	var columnCells int = rowCells

	const stateSize int = totalCells + 1
	var generation int = 0
	var preState = new([stateSize]string)
	var postState = new([stateSize]string)

	postState, generation = bigBang(postState, generation, totalCells)
	displayGrid(postState, generation, totalCells, rowCells, columnCells)

	//while true
	generation = generation + 1
	for i := 1; i <= totalCells; i++ {
		postState[i] = isAlive(preState[i])
	}
	for i := 1; i <= totalCells; i++ {
		preState[i] = postState[i]
	}
	displayGrid(postState, generation, totalCells, rowCells, columnCells)
}
