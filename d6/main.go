package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// //go:embed input_test.txt
// var input string

//go:embed input.txt
var input string

func main() {
	s := strings.Split(input, ",")
	var school = make([]int, 9, 9)

	for _, fish := range s {
		school[toInt(fish)]++
	}
	for i := 1; i <= 256; i++ {
		eggs := school[0]
		school[0] = school[1]
		school[1] = school[2]
		school[2] = school[3]
		school[3] = school[4]
		school[4] = school[5]
		school[5] = school[6]
		school[6] = school[7] + eggs
		school[7] = school[8]
		school[8] = eggs
	}
	totalfish := 0
	for _, fish := range school {
		totalfish += fish
	}
	fmt.Printf("School: %d\n", totalfish)
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
