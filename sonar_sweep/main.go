package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := readLines("data.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(determineIncrease(input))
	fmt.Println(threeMeasurementWindowSlice(input))
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, stringToInt(scanner.Text()))
	}

	return lines, scanner.Err()
}

func determineIncrease(slice []int) int {
	counter := 0

	for index, element := range slice {
		if index == 0 {
		} else if element > slice[index-1] {
			counter++
		}
	}

	return counter
}

func threeMeasurementWindowSlice(slice []int) int {
	measurementSlice := []int{}

	for i := 0; i <= len(slice)-3; i++ {
		sumOfASlice := addSlice(slice[i : i+3])

		measurementSlice = append(measurementSlice, sumOfASlice)
	}

	return determineIncrease(measurementSlice)
}

func addSlice(slice []int) int {
	sum := 0

	for _, v := range slice {
		sum += v
	}

	return sum
}

func stringToInt(text string) int {
	n, err := strconv.Atoi(text)

	if err != nil {
		fmt.Println(err)
	}

	return n
}
