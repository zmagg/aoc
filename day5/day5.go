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
	filename := "/home/zmagg/dev/aoc/day5/input"
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

	// Parse input:
	// create [][] slice 0->1000
	// if x!=x && y!=y, discard
	// else, place 1s where the line segment is
	// if a value is already there ++

	inputSlice := [][]int{}

	for i := 0; i < 1000; i++ {
		slice := []int{}
		for j := 0; j < 1000; j++ {
			slice = append(slice, 0)
		}
		inputSlice = append(inputSlice, slice)
	}

	for _, line := range input {
		res := strings.Split(line, "->")
		first := strings.Split(strings.TrimSpace(res[0]), ",")
		second := strings.Split(strings.TrimSpace(res[1]), ",")
		firstx, _ := strconv.Atoi(first[0])
		firsty, _ := strconv.Atoi(first[1])
		secondx, _ := strconv.Atoi(second[0])
		secondy, _ := strconv.Atoi(second[1])

		if firstx != secondx && firsty != secondy {
			if firstx > secondx && firsty > secondy {
				j := firsty
				for i := firstx; i >= secondx; i-- {
					inputSlice[i][j] = inputSlice[i][j] + 1
					j--
				}
				// 0,4 => 2,8
				//
			} else if firstx < secondx && firsty > secondy {
				j := firsty
				for i := firstx; i <= secondx; i++ {
					inputSlice[i][j] = inputSlice[i][j] + 1
					j--
				}
			} else if firstx < secondx && firsty < secondy {
				j := firsty
				for i := firstx; i <= secondx; i++ {
					inputSlice[i][j] = inputSlice[i][j] + 1
					j++
				}
			} else if firstx > secondx && firsty < secondy {
				j := firsty
				for i := firstx; i >= secondx; i-- {
					inputSlice[i][j] = inputSlice[i][j] + 1
					j++
				}
			}
		}

		if firstx == secondx {
			if firsty > secondy {

				for i := secondy; i <= firsty; i++ {
					inputSlice[firstx][i] = inputSlice[firstx][i] + 1
				}
			} else {
				for i := firsty; i <= secondy; i++ {
					inputSlice[firstx][i] = inputSlice[firstx][i] + 1
				}
			}
		}

		if firsty == secondy {
			if firstx > secondx {
				for i := secondx; i <= firstx; i++ {
					inputSlice[i][firsty] = inputSlice[i][firsty] + 1
				}
			} else {
				for i := firstx; i <= secondx; i++ {
					inputSlice[i][firsty] = inputSlice[i][firsty] + 1
				}
			}
		}

	}
	fmt.Print(computeOverlap(inputSlice))
}

func computeOverlap(input [][]int) int {
	counter := 0
	for _, i := range input {
		for _, j := range i {
			if j >= 2 {
				counter++
			}
		}
	}
	return counter
}

func printArray(input [][]int) {
	for _, i := range input {
		fmt.Println(i)
	}
}
