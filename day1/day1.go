package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "/home/zmagg/dev/aoc/day1/input"
	
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numIncreases :=0
	iterations := 0
	slide :=0
	first :=0 //pointer to first int in the sliding area
	second :=0 // pointer to second int
	third := 0
	for scanner.Scan() {

		val,err := strconv.Atoi(scanner.Text())
		
		if err == nil {
		
			if (iterations < 3) {
			if (iterations == 0 ) {
				first = val
			}
			if (iterations == 1) {
				second = val
			}
			if (iterations == 2) {
				third = val
			}
				slide = slide + val
			}

			if (iterations >=3) {
				// do slide comparisions
				newSlide := slide + val - first
				if (newSlide > slide) {
					numIncreases +=1
				}
				// move pointers around
				first = second
				second = third
				third = val
				slide = newSlide
			}
		}
		iterations +=1
	}

	fmt.Printf("number of increases %d", numIncreases)

}

