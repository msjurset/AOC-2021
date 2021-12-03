package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	depth, hpos, aim := 0, 0, 0
	for _, instruction := range getInstructions() {
		direction := instruction[0]
		coord := toInt(instruction[1])

		switch direction {
		case "forward":
			hpos += coord
			depth += aim * coord
		case "up":
			aim -= coord
		case "down":
			aim += coord
		}

	}
	fmt.Printf("Horizontal position: %d\n", hpos)
	fmt.Printf("Aim: %d\n", aim)
	fmt.Printf("Depth: %d\n\n", depth)

	fmt.Printf("depth * hpos: %d\n", depth*hpos)
}

func getInstructions() [][]string {
	fptr := flag.String("fpath", "input.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	check(err)

	defer func() {
		check(f.Close())
	}()

	s := bufio.NewScanner(f)

	var instructions [][]string
	for s.Scan() {
		check(s.Err())

		inst := strings.Split(s.Text(), " ")
		instructions = append(instructions, inst)
	}
	return instructions
}

func check(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}
