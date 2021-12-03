package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := readLines("data.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(divingPos(input))
}

func divingPos(input []string) int {
	x, y, z := 0, 0, 0
	direction := ""
	nCtr := 0

	for _, element := range input {
		direction, nCtr = split(element)

		if direction == "forward" {
			x += nCtr

			if z != 0 {
				y += z * nCtr
			}
		} else if direction == "up" {
			//y -= nCtr
			z -= nCtr
		} else if direction == "down" {
			//y += nCtr
			z += nCtr
		}
	}

	return x * y
}

func stringToInt(text string) int {
	n, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println(err)
	}

	return n
}

func split(text string) (string, int) {
	x := strings.Split(text, " ")

	return x[0], stringToInt(x[1])
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
