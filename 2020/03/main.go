package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	Field := loadInput()

	var items items
	items.tree = "#"

	var directions1 mover
	directions1.down = 1
	directions1.right = 1

	var directions2 mover
	directions2.down = 1
	directions2.right = 3

	var directions3 mover
	directions3.down = 1
	directions3.right = 5

	var directions4 mover
	directions4.down = 1
	directions4.right = 7

	var directions5 mover
	directions5.down = 2
	directions5.right = 1

	TreesSlopes1 := treeCounter(items, directions1, Field)
	fmt.Printf("\nEncountered %d trees in Slope 1\n", TreesSlopes1)

	TreesSlopes2 := treeCounter(items, directions2, Field)
	fmt.Printf("Encountered %d trees in Slope 2\n", TreesSlopes2)

	TreesSlopes3 := treeCounter(items, directions3, Field)
	fmt.Printf("Encountered %d trees in Slope 3\n", TreesSlopes3)

	TreesSlopes4 := treeCounter(items, directions4, Field)
	fmt.Printf("Encountered %d trees in Slope 4\n", TreesSlopes4)

	TreesSlopes5 := treeCounter(items, directions5, Field)
	fmt.Printf("Encountered %d trees in Slope 5\n", TreesSlopes5)

	fmt.Printf("---------------------------------\n            %d\n\n", TreesSlopes1*TreesSlopes2*TreesSlopes3*TreesSlopes4*TreesSlopes5)

}

func treeCounter(items items, directions mover, Field [323][31]string) int {
	var currentPos position
	currentPos.column = 0
	currentPos.row = 0

	numTrees := 0
	for i, y := range Field {
		for n, x := range y {
			//fmt.Println(i, n)
			if checkPosition(n, i, currentPos) {
				if x == items.tree {
					numTrees++
				}
				currentPos = calcNextPos(currentPos, directions)
				//fmt.Println(currentPos.row, currentPos.column, currentPos.slice, x)
			}

		}
	}
	return numTrees
}

func checkPosition(x int, y int, currentPos position) bool {
	//fmt.Println("check ", y, x, " - ", currentPos.row, currentPos.column, currentPos.slice)
	if y == currentPos.row {
		if x == currentPos.column {
			return true
		}
	}
	return false
}

func calcNextPos(currentPos position, directions mover) position {
	overlap := 31

	if directions.up > 0 {
		currentPos.row = currentPos.row - directions.up
	}

	if directions.down > 0 {
		currentPos.row = currentPos.row + directions.down
	}

	if directions.left > 0 {
		currentPos.column = currentPos.column - directions.left
		if currentPos.column < 0 {
			currentPos.column = currentPos.column + overlap
			currentPos.slice--
			if currentPos.slice < 0 {
				fmt.Errorf("Something wrong here")
			}
		}
	}

	if directions.right > 0 {
		currentPos.column = currentPos.column + directions.right

		if currentPos.column >= overlap {
			currentPos.column = currentPos.column - overlap
			currentPos.slice++
		}
	}

	return currentPos
}

type items struct {
	tree string
}

type position struct {
	column int
	row    int
	slice  int
}
type mover struct {
	down  int
	up    int
	left  int
	right int
}

//Todo: there must be a better way then to define the array so early
func loadInput() [323][31]string {
	dat, err := ioutil.ReadFile("input.txt")
	check(err)
	var Map [323][31]string
	lines := strings.Split(string(dat), "\n")
	for i, row := range lines {
		for n, column := range row {
			Map[i][n] = string(column)
		}
	}
	return Map
}
