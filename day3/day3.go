package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	filename := "/home/zmagg/dev/aoc/day3/input"
	bits := 12
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

	oxyOutput := input
	co2Output := input
	for i := 0; i < bits; i++ {
		if len(oxyOutput) > 1 {
			oxyOutput = saveOxyBitMatch(oxyOutput, i)
		}
		if len(co2Output) > 1 {
			co2Output = saveco2BitMatch(co2Output, i)
		}
	}

	oxygen := oxyOutput[0]
	co2 := co2Output[0]
	fmt.Printf("oxygen %f \n", convertBitStringToDecimal(oxygen, bits))
	fmt.Printf("co2 %f \n", convertBitStringToDecimal(co2, bits))

	fmt.Printf("life support rating %f", convertBitStringToDecimal(oxygen, bits)*convertBitStringToDecimal(co2, bits))

}

func convertBitStringToDecimal(input string, bits int) float64 {
	intVal := float64(0)
	for i := 0; i < bits; i++ {
		if input[i] == '1' {
			intVal += math.Pow(2, float64(bits-1-i))
		}
	}
	return intVal
}

func saveOxyBitMatch(input []string, index int) []string {
	common := 0
	numbLines := 0
	fmt.Printf("index %d %d ", index, len(input))
	for _, v := range input {

		val := v[index : index+1]
		if val == "1" {
			common++
		}
		numbLines++
	}

	fmt.Printf("common %d", common)
	fmt.Printf("numbLines %d\n", numbLines)

	output := []string{}

	commonZero := numbLines - common

	for _, v := range input {
		if commonZero <= common && (v[index:index+1] == "1") {
			output = append(output, v)
		} else if commonZero > common && (v[index:index+1] == "0") {
			output = append(output, v)
		}
	}

	if len(output) == 0 {
		return []string{input[len(input)-1]}
	}
	return output
}

func saveco2BitMatch(input []string, index int) []string {
	common := 0
	numbLines := 0
	for _, v := range input {
		val := v[index : index+1]
		if val == "1" {
			common++
		}
		numbLines++
	}

	output := []string{}

	commonZero := numbLines - common

	for _, v := range input {
		if commonZero > common && (v[index:index+1] == "1") {
			output = append(output, v)
		} else if commonZero <= common && (v[index:index+1] == "0") {
			output = append(output, v)
		}
	}

	return output
}
