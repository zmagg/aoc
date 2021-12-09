package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "/home/zmagg/dev/aoc/day9/input"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := []string{}
	length := 0
	for scanner.Scan() {
		input = append(input, scanner.Text())
		length++
	}

	mapp := [][]int{}

	for _, i := range input {
		row := []int{}
		for _, char := range i {
			charint := int(char - '0')
			row = append(row, charint)
		}
		mapp = append(mapp, row)
	}

	lowPoints := [][]int{}
	for i := 0; i < length; i++ {
		for j := 0; j < len(mapp[0]); j++ {
			if i != 0 {
				left := mapp[i-1][j]
				if left < mapp[i][j] {
					continue
				}
			}
			if j != 0 {
				up := mapp[i][j-1]
				if up < mapp[i][j] {
					continue
				}
			}

			if j != len(mapp[0])-1 {
				down := mapp[i][j+1]
				if down < mapp[i][j] {
					continue
				}
			}

			if i != length-1 {
				right := mapp[i+1][j]
				if right < mapp[i][j] {
					continue
				}
			}

			lowPoint := []int{i, j, mapp[i][j]}
			lowPoints = append(lowPoints, lowPoint)
		}
	}

	sum := 0
	for _, i := range lowPoints {
		sum = sum + i[2] + 1
	}

	fmt.Print(sum)
}
