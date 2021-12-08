package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board [][]map[int]int
type Set []int

func main() {
	filename := "/home/zmagg/dev/aoc/day4/input"
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

	boards := parseBoard(input)
	// displayBoard(boards)
	boards = playBingo(boards, input[0])

	highScore := 0
	worstBoard := make(Board, 0, 0)
	for _, board := range boards {

		if len(board) == 0 {
			continue
		}
		score := scoreBingo(board)
		if score > highScore {
			highScore = score
			worstBoard = board
		}
	}

	boardScore := 0
	numbJustCalled := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			mapp := worstBoard[i][j]
			for k, v := range mapp {
				if v > highScore && v != 1000 {
					boardScore += k
				}
				if v == highScore {
					numbJustCalled = k
				}
			}
		}
	}

	fmt.Print(numbJustCalled * boardScore)

}

func scoreBingo(board Board) int {
	// Score rows
	bestSet := []int{}
	bestPrintSet := []int{}

	for _, j := range board {
		noBingo := false
		rowScore := 0
		set := []int{}
		printSet := []int{}
		for _, mapp := range j {
			for k, v := range mapp {
				if v == 1000 {
					noBingo = true
					break
				}
				rowScore += v
				set = append(set, v)
				printSet = append(printSet, k)
			}
		}

		if (len(bestSet) == 0 || findLargest(bestSet) > findLargest(set)) && !noBingo {
			bestSet = set
			bestPrintSet = printSet
		}
	}

	// Score columns
	for i := 0; i < 5; i++ {
		noBingo := false
		set := []int{}
		printSet := []int{}
		for j := 0; j < 5; j++ {
			mapp := board[j][i]
			for k, v := range mapp {
				if v == 1000 {
					noBingo = true
					break
				}
				set = append(set, v)
				printSet = append(printSet, k)
			}
		}
		if (len(bestSet) == 0 || findLargest(bestSet) > findLargest(set)) && !noBingo {
			bestSet = set
			bestPrintSet = printSet
		}
	}

	fmt.Print(bestPrintSet)

	return findLargest(bestSet)

}

func findLargest(set []int) int {
	// Find largest index in bestSet

	bigVal := 0
	for _, i := range set {
		if i > bigVal {
			bigVal = i
		}
	}

	return bigVal
}

func playBingo(boards []Board, bingo string) []Board {

	bingoMap := map[int]int{}

	for index, numb := range strings.Split(bingo, ",") {
		intNumb, err := strconv.Atoi(numb)
		if err == nil {
			bingoMap[intNumb] = index
		}
	}

	for _, board := range boards {
		for _, j := range board {
			for _, mapp := range j {
				for k := range mapp {
					if index, ok := bingoMap[k]; ok {
						mapp[k] = index
					}
				}
			}
		}
	}
	return boards
}

// For testing
func displayBoard(boards []Board) {
	for _, board := range boards {
		for _, j := range board {
			for _, mapp := range j {
				for val, _ := range mapp {
					fmt.Print(val)
				}
				fmt.Printf(" ")
			}
			fmt.Printf("\n")
		}
		fmt.Printf("*****\n")
	}
}

func parseBoard(input []string) []Board {
	// Skip the first two lines as they don't express the board
	board := make(Board, 0, 0)
	returnBoards := make([]Board, 0, 0)

	boardSize := 5

	for i := 2; i < len(input); i++ {
		if ((i - 1) % (boardSize + 1)) == 0 {
			// empty line between boards
			continue
		}

		line := strings.Split(input[i], " ")
		hor := []map[int]int{}
		for _, val := range line {
			intval, err := strconv.Atoi(val)
			if err == nil {
				m := map[int]int{intval: 1000}
				hor = append(hor, m)
			}
		}
		board = append(board, hor)

		if (i % (boardSize + 1)) == 0 {
			returnBoards = append(returnBoards, board)
			board = make(Board, 0, 0)
		}
	}
	return returnBoards
}
