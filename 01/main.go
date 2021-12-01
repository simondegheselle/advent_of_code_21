package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	solvePart2(file)
}

func solvePart1(file *os.File) {
	scanner := bufio.NewScanner(file)

	amountOfIncrease := 0
	lastKnownVal := math.MaxInt32
	for scanner.Scan() {
		readVal, _ := strconv.Atoi(scanner.Text())
		if readVal > lastKnownVal {
			amountOfIncrease += 1
		}
		lastKnownVal = readVal
	}
	fmt.Println(amountOfIncrease)
}

func sum(numbs [3]int) int {
	total := 0
	for _, number := range numbs {
		total = total + number
	}
	return total
}

func solvePart2(file *os.File) {
	scanner := bufio.NewScanner(file)

	amountOfIncrease := 0

	history := [3]int{math.MaxInt32, 0, 0}
	i := 0
	iter := 0

	for scanner.Scan() {

		readVal, _ := strconv.Atoi(scanner.Text())

		old := sum(history)
		history[i] = readVal
		i = (i + 1) % 3
		new := sum(history)

		if new > old && iter >= 3 {
			amountOfIncrease += 1
		}
		iter += 1
	}
	fmt.Println(amountOfIncrease)
}
