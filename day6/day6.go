package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "/home/zmagg/dev/aoc/day6/input"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// Index corresponds to what day in their cycle they are in
	fishes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	inputString := strings.Split(input[0], ",")

	for _, i := range inputString {
		day, _ := strconv.Atoi(i)
		fishes[day]++
	}

	for i := 0; i < 256; i++ {
		newFishes := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for day := 0; day <= 5; day++ {
			newFishes[day] = fishes[day+1]
		}

		newFishes[8] = fishes[0]
		newFishes[6] = fishes[7] + fishes[0]
		newFishes[7] = fishes[8]

		fishes = newFishes
	}

	numbFishes := 0
	for _, day := range fishes {
		numbFishes = numbFishes + day
	}

	fmt.Print(numbFishes)
}
