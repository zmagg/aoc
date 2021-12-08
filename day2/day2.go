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
	filename := "/home/zmagg/dev/aoc/day2/input"

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	vert := 0
	hor := 0
	aim := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		direction := s[0]
		count, err := strconv.Atoi(s[1])
		if err == nil {
			if direction == "forward" {
				hor += count
				vert += aim * count

			} else if direction == "down" {
				aim += count

			} else if direction == "up" {
				aim -= count
			}
		}
	}

	fmt.Printf("position %d", vert*hor)

}
