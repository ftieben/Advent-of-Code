package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intToIntArray(input int) []int {
	tmp := strconv.Itoa(input)
	tmpArrayOfLetters := strings.Split(tmp, "")

	var cups []int
	for _, value := range tmpArrayOfLetters {
		_tmp, err := strconv.Atoi(value)
		check(err)
		cups = append(cups, _tmp)
	}
	return cups
}

func intArrayToInt(input []int) int {
	targetString := ""
	for _, value := range input {
		targetString = targetString + strconv.Itoa(value)
	}
	targetInt, _ := strconv.Atoi(targetString)
	return targetInt
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// -------------------------------------------------------------------

func main() {
	input := 784235916
	play(intToIntArray(input))
}

func play(cups []int) []int {
	/*
		1. The crab picks up the three cups that are immediately clockwise of the current cup. They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.

		2. The crab selects a destination cup: the cup with a label equal to the current cup's label minus one. If this would select one of the cups that was just picked up,
		the crab will keep subtracting one until it finds a cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's label,
		it wraps around to the highest value on any cup's label instead.

		3. The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.

		4. The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
	*/

	// 1. -------------------------------------------------------------------- //
	currentPos := 0
	fmt.Println("Position:", currentPos, "Cups:", cups)
	numPigs := 3
	var pigs []int
	pigged := 0
	for i := range cups {
		i = i - pigged
		if i == currentPos+1 && pigged < numPigs {
			pigs = append(pigs, cups[i])
			cups = remove(cups, i)
			pigged++
		}
	}

	// 2. -------------------------------------------------------------------- //
	destination := calcDestination(currentPos, cups)
	fmt.Println("Position:", currentPos, "Cups:", cups, "Destination:", destination, "Pickups:", pigs)

	return cups
}

func calcDestination(currentPos int, cups []int) int {
	currentLabel := 0
	for i, cup := range cups {
		if i == currentPos {
			currentLabel = cup
		}
	}

	nextLabel := currentLabel - 1
	nextPos := 0
	for i, cup := range cups {
		if nextLabel == cup {
			nextPos = i
		}
	}
	return nextPos
}
