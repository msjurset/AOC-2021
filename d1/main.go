package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fptr := flag.String("fpath", "input.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	acc := 0
	a, b, c, d := 0, 0, 0, 0

	// read the first line
	readNext(s)

	n := getNum(err, s)
	a = n
	readNext(s)

	n = getNum(err, s)
	a += n
	b = n
	readNext(s)

	n = getNum(err, s)
	a += n
	b += n
	c = n
	readNext(s)

	n = getNum(err, s)
	b += n
	c += n
	d = n

	if b > a {
		acc++
	}

	a = b
	b = c
	c = d

	for s.Scan() {
		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
		n = getNum(err, s)
		b += n
		c += n
		d = n

		if b > a {
			acc++
		}
		a = b
		b = c
		c = d
	}
	fmt.Println(acc)
}

func getNum(err error, s *bufio.Scanner) int {
	n, err := strconv.Atoi(s.Text())
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func readNext(s *bufio.Scanner) {
	s.Scan()
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
